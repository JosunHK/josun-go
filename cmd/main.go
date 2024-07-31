package main

import (
	"fmt"
	"os"

	"github.com/JosunHK/josun-go.git/cmd/cfg"
	"github.com/JosunHK/josun-go.git/cmd/database"
	"github.com/JosunHK/josun-go.git/cmd/handlers/api"
	"github.com/JosunHK/josun-go.git/cmd/handlers/pages"
	"github.com/JosunHK/josun-go.git/cmd/middleware"

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
	cfg, err := cfg.CfgInit()
	PORT := os.Getenv("PORT")
	if err != nil {
		log.Panic(err)
		return
	}

	log.Info("db cred ", os.Getenv("DB_CREDENTIALS"))
	if err := database.InitDB(os.Getenv("DB_CREDENTIALS")); err != nil {
		log.Panic(err)
		return
	}
	defer database.DB.Close()

	//static files
	e := echo.New()
	e.Static("/static", "web/static")

	//end points
	e.GET("/", middleware.Content(pages.Layout, cfg))

	//end points
	e.GET("/Users", middleware.Service(api.GetUsers, cfg))

	//exit ->
	e.Logger.Fatal(e.Start(PORT))
}
