package middleware

import (
	"github.com/JosunHK/josun-go.git/cmd/cfg"
	"github.com/JosunHK/josun-go.git/cmd/i18n"
	"github.com/JosunHK/josun-go.git/cmd/util"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

type contentHandler func(echo.Context, i18n.Transl) error
type serviceHandler func(echo.Context) (err error, statusCode int, resObj interface{})

func HTML(next contentHandler, cfg cfg.Cfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		locale := util.GetLocaleFromCookie(c)
		T := func(key string) string {
			return cfg.I18n.T(locale, key)
		}

		return next(c, T)
	}
}

func JSON(handler serviceHandler, cfg cfg.Cfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		err, statusCode, resObj := handler(c)
		if err != nil {
			log.Error(err)
			return c.String(statusCode, err.Error())
		}

		return c.JSON(statusCode, resObj)
	}
}
