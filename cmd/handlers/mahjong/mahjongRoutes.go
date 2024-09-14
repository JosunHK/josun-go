package mahjong

import (
	"github.com/JosunHK/josun-go.git/cmd/layout"
	"github.com/JosunHK/josun-go.git/cmd/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	//main endpoints
	e.GET("/", middleware.Pages(layout.Layout, RoomSelect))
	e.GET("/mahjong/room/:code", middleware.Pages(layout.Layout, Room))

	//api endpoints
	e.GET("/mahjong/room/create", middleware.HTML(RoomSetting))
	e.GET("/mahjong/result/:id", middleware.Pages(layout.Layout, GameResult))
	e.POST("/mahjong/room/create", middleware.Redirect(RoomCreate))
}
