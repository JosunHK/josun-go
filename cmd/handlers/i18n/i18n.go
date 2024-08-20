package i18n

import (
	i18nTemplates "github.com/JosunHK/josun-go.git/web/templates/contents/i18n"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Table(c echo.Context) templ.Component {
	locale := c.Param("locale")
	return i18nTemplates.I18n(locale)
}
