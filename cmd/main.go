package main

import (
	"fmt"
	"os"

	"github.com/JosunHK/josun-go.git/cmd/database"
	"github.com/JosunHK/josun-go.git/cmd/handlers/api"
	"github.com/JosunHK/josun-go.git/cmd/handlers/pages"
	"github.com/JosunHK/josun-go.git/cmd/middleware"
	"github.com/JosunHK/josun-go.git/cmd/util/i18n"
	"github.com/JosunHK/josun-go.git/pkg/twmerge"

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

	//merger for tailwind
	config := twmerge.MakeDefaultConfig()
	_ = twmerge.CreateTwMerge(config, nil) // config, cache (if nil default will be used)
}

func main() {
	PORT := os.Getenv("PORT")
	if err := database.InitDB(os.Getenv("DB_CREDENTIALS")); err != nil {
		log.Panic(err)
		return
	}
	if err := i18n.InitI18n(); err != nil {
		log.Panic(err)
		return
	}

	defer database.DB.Close()

	//static files
	e := echo.New()
	e.Static("/static", "web/static")

	//end points
	e.GET("/", middleware.HTML(pages.Layout))
	e.GET("/playground", middleware.HTML(pages.Layout))

	//dummy api
	e.GET("/Users", middleware.JSON(api.GetUsers))
	e.POST("/Users", middleware.JSON(api.PostUser))

	//exit ->
	e.Logger.Fatal(e.Start(PORT))
}
