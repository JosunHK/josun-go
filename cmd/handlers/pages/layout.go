package pages

import (
	"github.com/JosunHK/josun-go.git/cmd/i18n"
	"github.com/JosunHK/josun-go.git/cmd/util"
	"github.com/JosunHK/josun-go.git/web/templates/layout"
	"github.com/labstack/echo/v4"
)

func Layout(c echo.Context, T i18n.Transl) error {
	return util.HTML(c, layout.Layout("hello_world", T))
}
