package layout

import (
	responseUtil "github.com/JosunHK/josun-go.git/cmd/util/response"
	layoutTemplates "github.com/JosunHK/josun-go.git/web/templates/layout"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Layout(c echo.Context, content templ.Component) error {
	return responseUtil.HTML(c, layoutTemplates.Layout(content))
}

func Component(c echo.Context, content templ.Component) error {
	return responseUtil.HTML(c, content)
}
