package responseUtil

import (
	"context"

	i18nUtil "github.com/JosunHK/josun-go.git/cmd/util/i18n"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func HTML(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(overrideAndGetContext(c), c.Response().Writer)
}

func overrideAndGetContext(c echo.Context) context.Context {
	context := c.Request().Context()
	context = overrideContextWithLocale(c)
	return context
}

func overrideContextWithLocale(c echo.Context) context.Context {
	var locale string
	cookie, err := c.Cookie(i18nUtil.LOCALE_SETTING_ID)
	if err != nil {
		locale = "en"
	} else {
		locale = cookie.Value
	}

	return context.WithValue(c.Request().Context(), i18nUtil.LOCALE_SETTING_ID, locale)
}
