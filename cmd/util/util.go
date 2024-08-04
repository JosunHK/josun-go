package util

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func HTML(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

// dummy for now
func GetLocaleFromCookie(c echo.Context) string {
	return "en"
}

func GetIsDarkFromCookie(c echo.Context) bool {
	cookie, err := c.Cookie("username")
	if err != nil {
		log.Error("Cookie not found", err)
		return true // default dark mode
	}

	val, err := strconv.ParseBool(cookie.Value)
	if err != nil {
		log.Error("Cookie value not a bool", err)
		return true // default dark mode
	}

	return val
}
