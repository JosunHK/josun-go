package middleware

import (
	"github.com/JosunHK/josun-go.git/src/cfg"
	"github.com/JosunHK/josun-go.git/src/i18n"
	"github.com/JosunHK/josun-go.git/src/util"
	"github.com/labstack/echo/v4"
)

func Content(next func(echo.Context, i18n.Transl) error, cfg cfg.Cfg) echo.HandlerFunc {

	return func(c echo.Context) error {
		locale := util.GetLocaleFromCookie(c)

		T := func(key string) string {
			return cfg.I18n.T(locale, key)
		}

		return next(c, T)
	}
}
