package mahjong

import (
	"github.com/JosunHK/josun-go.git/cmd/util"
	mahjongTemplates "github.com/JosunHK/josun-go.git/web/templates/contents/mahjong"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RoomSelect(c echo.Context) templ.Component {
	return mahjongTemplates.RoomSelect()
}

func RoomCreate(c echo.Context) error {
	return util.HTML(c, mahjongTemplates.RoomCreate())
}

func Test(c echo.Context) templ.Component {
	return mahjongTemplates.Test()
}
