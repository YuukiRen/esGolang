package domain

import (
	"context"
	"errors"
	"github.com/labstack/echo"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
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
func IsRequestValid(request interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return false, err
	}
	return true, nil
}

func CreateIndexIfDoesNotExist(ctx context.Context, client *elastic.Client, indexName string) error {
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	res, err := client.CreateIndex(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if !res.Acknowledged {
		return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct")
	}
	return nil
}