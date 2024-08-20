package user

import (
	"github.com/JosunHK/josun-go.git/cmd/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	//main endpoints
	e.GET("/Users", middleware.JSON(GetUsers))
	e.POST("/Users", middleware.JSON(PostUser))
}
