package util

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"runtime"
)

const wrapFormat = `{"file":"%s:%d", "func":"%s", "error":"%s"}`

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	pc, file, line, _ := runtime.Caller(1)
	return fmt.Errorf(wrapFormat, file, line, runtime.FuncForPC(pc).Name(), err)
}

func LogErrWithReqId(c *fiber.Ctx, err error) {
	slog.Error(
		"internal", "request-id", c.Context().UserValue("request-id"), "error", err)
}
