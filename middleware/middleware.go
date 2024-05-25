package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
		userId := c.Context().UserValue("user-id")
		if userId == nil {
			userId = "not-logged-in"
		}
		baseAttributes := []slog.Attr{
			slog.String("user-id", fmt.Sprintf("%v", userId)),
			slog.String("request-id", requestID),
			slog.String("method", string(c.Context().Method())),
			slog.Int("status", status),
			slog.String("path", c.Path()),
			slog.String("query", string(c.Request().URI().QueryString())),
			slog.String("ip", c.Context().RemoteIP().String()),
			slog.Any("x-forwarded-for", c.IPs()),
			slog.String("referer", c.Get(fiber.HeaderReferer)),
			// slog.String("user-agent", string(c.Context().UserAgent())),
			slog.Any("body", string(c.Body())),
			slog.String("latency", fmt.Sprintf("%dms", end.Sub(start).Milliseconds())),
		}

		level := slog.LevelInfo
		if status >= http.StatusInternalServerError {
			level = slog.LevelWarn
		} else if status >= http.StatusBadRequest && status < http.StatusInternalServerError {
			level = slog.LevelError
		}
		logger.LogAttrs(c.UserContext(), level, "request", baseAttributes...)

		return err
	}
}

func NewNoSigninUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if sessId := c.Cookies("session_id"); sessId != "" {
			return c.Redirect("/")
		}
		return c.Next()
	}
}

func NewSessionAuth(redisClient *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessId := c.Cookies("session_id")
		userId, err := redisClient.Get(context.Background(), sessId).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				return redirectByURL(c)
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		sessIds, err := redisClient.LRange(context.Background(), userId, 0, -1).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				return redirectByURL(c)
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		for _, sid := range sessIds {
			if sessId == sid {
				userIdInt, _ := strconv.Atoi(userId)
				c.Context().SetUserValue("user-id", userIdInt)

				return c.Next()
			}
		}

		return redirectByURL(c)
	}
}

func redirectByURL(c *fiber.Ctx) error {
	if strings.HasPrefix(c.Path(), "/api") {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	redirectURL := "/signin?next=" + url.QueryEscape(c.OriginalURL())
	return c.Redirect(redirectURL)
}
