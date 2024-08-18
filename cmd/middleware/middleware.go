package middleware

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

type requestHandler func(echo.Context) error
type pageHandler func(echo.Context, templ.Component) error
type serviceHandler func(echo.Context) (err error, statusCode int, resObj interface{})
type PageHandler func(echo.Context) templ.Component

func StaticPages(next pageHandler, content templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c, content)
	}
}

func Pages(next pageHandler, p PageHandler) echo.HandlerFunc {
	return func(c echo.Context) error {

		return next(c, p(c))
	}
}

func HTML(next requestHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func JSON(handler serviceHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		err, statusCode, resObj := handler(c)
		if err != nil {
			log.Error(err)
			return c.String(statusCode, err.Error())
		}

		return c.JSON(statusCode, resObj)
	}
}
