package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// bodyWriter 用于捕获响应体
type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 统一打印请求/响应日志
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := uuid.New().String()
		c.Set("RequestID", requestID)
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		logrus.WithFields(logrus.Fields{
			"request_id": requestID,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"query":      c.Request.URL.RawQuery,
			"body":       string(bodyBytes),
		}).Info("请求开始")

		respBody := &bytes.Buffer{}
		writer := &bodyWriter{
			ResponseWriter: c.Writer,
			body:           respBody,
		}
		c.Writer = writer

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		fields := logrus.Fields{
			"requestId": requestID,
			"status":    status,
			"latencyMs": latency.Milliseconds(),
			"response":  respBody.String(),
		}

		if len(c.Errors) > 0 {
			errMsgs := ""
			for _, e := range c.Errors {
				errMsgs += e.Error() + " | "
			}
			fields["error"] = errMsgs
		}

		if status >= http.StatusBadRequest {
			logrus.WithFields(fields).Error("请求失败")
		} else {
			logrus.WithFields(fields).Info("请求结束")
		}
	}
}
