package middleware

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/ilcm96/dku-aegis-library/util"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ilcm96/dku-aegis-library/model"
	"github.com/redis/go-redis/v9"

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
		data, err := redisClient.Get(context.Background(), sessId).Bytes()
		// 조회되지 않는다면 redis.Nil(키 값을 찾을 수 없음)이면 리다이렉트, 이외 에러면 500을 반환한다
		if err != nil {
			if errors.Is(err, redis.Nil) {
				return redirectToSignInURL(c)
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		sess, err := decodeSessionData(data)
		if err != nil {
			util.LogErrWithReqId(c, err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// userId 를 키 값으로 가지는 리스트를 가져온다
		sessIds, err := redisClient.LRange(context.Background(), strconv.Itoa(sess.UserId), 0, -1).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				return redirectToSignInURL(c)
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// 조회된 리스트에 쿠키에서 가져온 세션 값이 없다면 리다이렉트
		// 일반적으로 이 상황은 다중 로그인 상황에서 탈퇴를 했을 때
		// 탈퇴 버튼을 누르지 않은 클라이언트에서 접속을 시도할 때 발생한다
		if isSessIdInList(sessId, sessIds) {
			c.Context().SetUserValue("user-id", sess.UserId)
			c.Context().SetUserValue("is-admin", sess.IsAdmin)
			return c.Next()
		}

		return redirectToSignInURL(c)
	}
}

func NewRenewSession(redisClient *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessId := c.Cookies("session_id")

		ttl, err := redisClient.TTL(context.Background(), sessId).Result()
		if ttl < 5*time.Minute {
			if err = redisClient.Expire(context.Background(), sessId, 10*time.Minute).Err(); err != nil {
				util.LogErrWithReqId(c, err)
				return c.SendStatus(fiber.StatusInternalServerError)
			}

			c.Cookie(&fiber.Cookie{
				Name:     "session_id",
				Value:    sessId,
				Path:     "/",
				Expires:  time.Now().Add(10 * time.Minute),
				HTTPOnly: true,
				SameSite: "Lax",
			})
			return c.Next()
		}

		return c.Next()
	}
}

func NewIsAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !c.Context().UserValue("is-admin").(bool) {
			if isApiURL(c.Path()) {
				return c.SendStatus(fiber.StatusForbidden)
			}
			return c.Redirect("/")
		}
		return c.Next()
	}
}

func decodeSessionData(data []byte) (model.Session, error) {
	buf := bytes.NewBuffer(data)
	var sess model.Session

	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&sess); err != nil {
		return sess, err
	}

	return sess, nil
}

func isSessIdInList(sessId string, list []string) bool {
	for _, id := range list {
		if sessId == id {
			return true
		}
	}
	return false
}

func redirectToSignInURL(c *fiber.Ctx) error {
	if isApiURL(c.Path()) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	redirectURL := "/signin?next=" + url.QueryEscape(c.OriginalURL())
	return c.Redirect(redirectURL)
}

func isApiURL(path string) bool {
	if strings.HasPrefix(path, "/api") {
		return true
	}
	return false
}
