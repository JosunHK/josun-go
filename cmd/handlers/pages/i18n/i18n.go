package i18nContent

import (
	"github.com/JosunHK/josun-go.git/web/templates/contents"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Content(c echo.Context) templ.Component {
	locale := c.Param("locale")
	return contents.I18n(locale)
}
