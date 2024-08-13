package middleware

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

type contentHandler func(echo.Context, templ.Component) error
type serviceHandler func(echo.Context) (err error, statusCode int, resObj interface{})

func HTML(next contentHandler, content templ.Component) echo.HandlerFunc {
	log.Info("log is working")
	return func(c echo.Context) error {
		return next(c, content)
	}
}

func JSON(handler serviceHandler) echo.HandlerFunc {
	log.Info("log is working")
	return func(c echo.Context) error {
		err, statusCode, resObj := handler(c)
		if err != nil {
			log.Error(err)
			return c.String(statusCode, err.Error())
		}

		return c.JSON(statusCode, resObj)
	}
}
