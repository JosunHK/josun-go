package mahjong

import (
	"github.com/JosunHK/josun-go.git/cmd/layout"
	"github.com/JosunHK/josun-go.git/cmd/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	//main endpoints
	e.GET("/mahjong", middleware.Pages(layout.Layout, RoomSelect))
	e.GET("/mahjong/room/test", middleware.Pages(layout.Layout, Test))

	//api endpoints
	e.GET("/mahjong/room/create", middleware.HTML(RoomSetting))
	e.POST("/mahjong/room/create", middleware.HTML(RoomCreate))
}
