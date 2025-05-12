package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func LoggingMiddleware(logger zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			req := c.Request()
			res := c.Response()

			err := next(c)

			// リクエストの属性を取得
			latency := time.Since(start)
			statusCode := res.Status
			method := req.Method
			uri := req.RequestURI
			ip := c.RealIP()

			// zap で構造化ログを出力
			logger.Info("request",
				zap.String("method", method),
				zap.String("uri", uri),
				zap.Int("status", statusCode),
				zap.Duration("latency", latency),
				zap.String("ip", ip),
			)

			return err
		}
	}
}
