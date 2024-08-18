package main

import (
	"fmt"
	"os"

	"github.com/JosunHK/josun-go.git/cmd/database"
	"github.com/JosunHK/josun-go.git/cmd/handlers/api"
	i18nAPI "github.com/JosunHK/josun-go.git/cmd/handlers/api/i18n"
	"github.com/JosunHK/josun-go.git/cmd/handlers/pages"
	i18nContent "github.com/JosunHK/josun-go.git/cmd/handlers/pages/i18n"
	"github.com/JosunHK/josun-go.git/cmd/middleware"
	"github.com/JosunHK/josun-go.git/cmd/util/i18n"
	"github.com/JosunHK/josun-go.git/pkg/twmerge"
	"github.com/JosunHK/josun-go.git/web/templates/contents"
	"github.com/JosunHK/josun-go.git/web/templates/contents/mahjong"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

func init() {
	file, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
		log.SetLevel(log.DebugLevel)
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

	e.GET("/", middleware.StaticPages(pages.Layout, contents.Playground()))
	e.GET("/mahjong", middleware.StaticPages(pages.Layout, mahjong.RoomSelect()))
	e.GET("/mahjong/room/create", middleware.StaticPages(pages.Component, mahjong.RoomCreate()))
	e.GET("/i18n/:locale", middleware.Pages(pages.Layout, i18nContent.Content))

	//dummy api
	e.GET("/i18n/items/:locale", middleware.HTML(i18nAPI.GetItems))
	e.POST("/i18n/items/:locale", middleware.HTML(i18nAPI.AddItems))
	e.DELETE("/i18n/items/:locale", middleware.HTML(i18nAPI.DeleteItems))

	e.GET("/Users", middleware.JSON(api.GetUsers))
	e.POST("/Users", middleware.JSON(api.PostUser))

	//exit ->
	e.Logger.Fatal(e.Start(PORT))
}
