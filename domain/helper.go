package domain

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"time"
)

type JSONResponse struct {
	Data interface{}	`json:"data"`
	Message string		`json:"message"`
}
func MakeLogEntry(c echo.Context) *logrus.Entry {
	if c == nil {
		return logrus.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return logrus.WithFields(logrus.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}