package middleware

import (
	"context"
	"net/http"

	"github.com/JosunHK/josun-go.git/cmd/util/cookie"
	i18nUtil "github.com/JosunHK/josun-go.git/cmd/util/i18n"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

type requestHandler func(echo.Context) error
type redirectHandler func(echo.Context) (string, error)
type pageHandler func(echo.Context, templ.Component) error
type serviceHandler func(echo.Context) (err error, statusCode int, resObj interface{})
type PageHandler func(echo.Context) templ.Component

func StaticPages(next pageHandler, content templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie.ManageGuestSession(&c)
		return next(c, content)
	}
}

func Pages(next pageHandler, p PageHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie.ManageGuestSession(&c)
		return next(c, p(c))
	}
}

func HTML(next requestHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func Redirect(next redirectHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		url, err := next(c)
		if err != nil {
			return err
		}

		c.Response().Header().Add("hx-redirect", url)
		return c.NoContent(http.StatusCreated)
	}
}

func NoContent(next requestHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			log.Error(err)
			return c.String(500, err.Error())
		}

		return c.NoContent(200)
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

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			log.Error(err)
		}
		return err
	}
}

func WithLocale(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.SetRequest(c.Request().WithContext(overrideContextWithLocale(c)))
		err := next(c)
		if err != nil {
			log.Error(err)
		}
		return err
	}
}

func overrideContextWithLocale(c echo.Context) context.Context {
	var locale string
	cookie, err := c.Cookie(i18nUtil.LOCALE_SETTING_ID)
	if err != nil {
		locale = "zh"
	} else {
		locale = cookie.Value
	}

	return context.WithValue(c.Request().Context(), i18nUtil.LOCALE_SETTING_ID, locale)
}
