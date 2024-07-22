package main

import (
	"fmt"
	"os"

	"github.com/JosunHK/josun-go.git/src/cfg"
	"github.com/JosunHK/josun-go.git/src/i18n"
	"github.com/JosunHK/josun-go.git/src/middleware"
	"github.com/JosunHK/josun-go.git/src/util"
	"github.com/JosunHK/josun-go.git/templates"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg, err := cfg.CfgInit()
	if err != nil {
		fmt.Println(err)
		return
	}

	PORT := os.Getenv("PORT")

	e := echo.New()
	e.Static("/static", "static")

	e.GET("/", middleware.Content(layout, cfg))

	e.Logger.Fatal(e.Start(PORT))
}

func layout(c echo.Context, T i18n.Transl) error {
	return util.HTML(c, templates.Layout("hello_world", T))
}
