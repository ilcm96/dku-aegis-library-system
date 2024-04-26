package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewSlog(logger *slog.Logger) fiber.Handler {
	var (
		once       sync.Once
		errHandler fiber.ErrorHandler
	)

	return func(c *fiber.Ctx) error {
		once.Do(func() {
			errHandler = c.App().ErrorHandler
		})

		requestID := uuid.New().String()
		c.Context().SetUserValue("request-id", requestID)
		c.Set("X-Request-ID", requestID)

		start := time.Now()
		err := c.Next()
		if err != nil {
			if err := errHandler(c, err); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}
		end := time.Now()

		status := c.Response().StatusCode()
		baseAttributes := []slog.Attr{
			slog.String("method", string(c.Context().Method())),
			slog.Int("status", status),
			slog.String("path", c.Path()),
			slog.String("query", string(c.Request().URI().QueryString())),
			slog.String("route", c.Route().Path),
			slog.String("ip", c.Context().RemoteIP().String()),
			slog.Any("x-forwarded-for", c.IPs()),
			slog.String("referer", c.Get(fiber.HeaderReferer)),
			// slog.String("user-agent", string(c.Context().UserAgent())),
			// slog.Any("body", string(c.Body())),
			slog.String("latency", fmt.Sprintf("%dms", end.Sub(start).Milliseconds())),
		}

		level := slog.LevelInfo
		if status >= http.StatusInternalServerError {
			level = slog.LevelWarn
		} else if status >= http.StatusBadRequest && status < http.StatusInternalServerError {
			level = slog.LevelError
		}
		logger.LogAttrs(c.UserContext(), level, requestID, baseAttributes...)

		return err
	}
}

func NewJwt() fiber.Handler {
	return keyauth.New(keyauth.Config{
		KeyLookup: "cookie:token",
		Validator: validateJWT,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			redirectURL := "/login?next=" + url.QueryEscape(c.OriginalURL())
			return c.Redirect(redirectURL)
		},
	})
}

func validateJWT(c *fiber.Ctx, token string) (bool, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("jwt-secret"), nil
	})
	if err != nil || !parsedToken.Valid {
		return false, c.Redirect("/login")
	}
	return true, nil
}
