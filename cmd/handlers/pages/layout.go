package pages

import (
	"github.com/JosunHK/josun-go.git/cmd/util"
	"github.com/JosunHK/josun-go.git/web/templates/contents"
	"github.com/JosunHK/josun-go.git/web/templates/layout"
	"github.com/labstack/echo/v4"
)

func Layout(c echo.Context) error {
	return util.HTML(c, layout.Layout(contents.Playground()))
}

func PlayGround(c echo.Context) error {
	return util.HTML(c, layout.Layout(contents.Playground()))
}
