package pages

import (
	"github.com/JosunHK/josun-go.git/cmd/util"
	"github.com/JosunHK/josun-go.git/web/templates/layout"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Layout(c echo.Context, content templ.Component) error {
	return util.HTML(c, layout.Layout(content))
}

func Component(c echo.Context, content templ.Component) error {
	return util.HTML(c, content)
}
