package pages

import (
	"github.com/JosunHK/josun-go.git/src/i18n"
	"github.com/JosunHK/josun-go.git/src/util"
	"github.com/JosunHK/josun-go.git/templates"
	"github.com/labstack/echo/v4"
)

func Layout(c echo.Context, T i18n.Transl) error {
	return util.HTML(c, templates.Layout("hello_world", T))
}
