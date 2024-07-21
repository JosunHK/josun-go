package main

import (
	"os"

	"github.com/JosunHK/josun-go.git/src/util"
	"github.com/JosunHK/josun-go.git/templates"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	PORT := os.Getenv("PORT")

	e := echo.New()
	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return util.HTML(c, templates.Layout("bruh"))
	})

	e.Logger.Fatal(e.Start(PORT))
}
