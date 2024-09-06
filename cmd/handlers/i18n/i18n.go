package i18n

import (
	"net/http"
	"time"

	i18nUtil "github.com/JosunHK/josun-go.git/cmd/util/i18n"
	i18nTemplates "github.com/JosunHK/josun-go.git/web/templates/contents/i18n"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Table(c echo.Context) templ.Component {
	locale := c.Param("locale")
	return i18nTemplates.I18n(locale)
}

func SetLocale(c echo.Context) error {
	locale := c.Param("locale")
	cookie := new(http.Cookie)
	cookie.Name = i18nUtil.LOCALE_SETTING_ID
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Path = "/"
	cookie.Secure = true
	cookie.Value = locale
	cookie.Expires = time.Now().Add(240 * time.Hour) //10 days
	c.SetCookie(cookie)

	return nil
}
