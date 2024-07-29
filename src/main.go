package main

import (
	"fmt"
	"os"

	"github.com/JosunHK/josun-go.git/src/cfg"
	"github.com/JosunHK/josun-go.git/src/handlers/pages"
	"github.com/JosunHK/josun-go.git/src/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

func init() {
	file, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	err = godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	//init
	cfg, err := cfg.CfgInit()
	PORT := os.Getenv("PORT")
	if err != nil {
		fmt.Println(err)
		return
	}

	//static files
	e := echo.New()
	e.Static("/static", "static")

	//end points
	e.GET("/", middleware.Content(pages.Layout, cfg))

	//end points
	e.GET("/api", middleware.Content(pages.Layout, cfg))

	//exit ->
	e.Logger.Fatal(e.Start(PORT))
}
