package dummy

import (
	"github.com/JosunHK/josun-go.git/cmd/layout"
	"github.com/JosunHK/josun-go.git/cmd/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	//dummy routes
	e.GET("/dummy", middleware.Pages(layout.Layout, Dummy))

	e.GET("/dummy/odometer", middleware.HTML(Odometer))

	//e.GET("/dummy/odometer", Odometer)
}
