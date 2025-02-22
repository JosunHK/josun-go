package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/JosunHK/josun-go.git/cmd/database"
	"github.com/JosunHK/josun-go.git/cmd/handlers/dummy"
	"github.com/JosunHK/josun-go.git/cmd/handlers/i18n"
	"github.com/JosunHK/josun-go.git/cmd/handlers/mahjong"
	"github.com/JosunHK/josun-go.git/cmd/handlers/user"
	"github.com/JosunHK/josun-go.git/cmd/layout"
	"github.com/JosunHK/josun-go.git/cmd/middleware"
	"github.com/JosunHK/josun-go.git/cmd/pubsub"
	i18nUtil "github.com/JosunHK/josun-go.git/cmd/util/i18n"
	playgroundTemplates "github.com/JosunHK/josun-go.git/web/templates/contents/playground"
	twmerge "github.com/Oudwins/tailwind-merge-go/pkg/twmerge"
	eMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	BUILD := os.Getenv("BUILD")
	if BUILD == "dev" {
		file, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
			log.SetLevel(log.DebugLevel)
		} else {
			log.Info("Failed to log to file, using default stderr")
		}
	} else {
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)
	}

	//merger for tailwind
	config := twmerge.MakeDefaultConfig()
	_ = twmerge.CreateTwMerge(config, nil) // config, cache (if nil default will be used)
}

const MYSQL_PARAMS = "?parseTime=true&loc=Local"

func main() {
	PORT := os.Getenv("PORT")
	PORT = ":" + PORT
	if err := database.InitDB(os.Getenv("DB_CREDENTIALS") + MYSQL_PARAMS); err != nil {
		log.Error(err)
		return
	}
	if err := i18nUtil.InitI18n(); err != nil {
		log.Error(err)
		return
	}

	defer database.DB.Close()

	routers, err := pubsub.NewRouters()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(eMiddleware.Recover())
	e.Use(middleware.Logger)
	e.Use(eMiddleware.Logger())
	e.Use(middleware.WithLocale)
	e.Pre(eMiddleware.RemoveTrailingSlashWithConfig(
		eMiddleware.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		},
	))

	//static files
	e.Static("/static", "web/static")
	e.File("/favicon.ico", "web/static/favicon.ico")

	e.GET("/playground", middleware.StaticPages(layout.Layout, playgroundTemplates.Playground()))
	go pubsub.StartEventsRouter(context.Background(), routers)
	go pubsub.StartSSERouter(context.Background(), routers)

	pubsub.NewHandler(e, routers.EventBus, routers.SSERouter)
	i18n.RegisterRoutes(e)
	mahjong.RegisterRoutes(e)
	user.RegisterRoutes(e)
	dummy.RegisterRoutes(e)

	//exit ->
	e.Logger.Fatal(e.Start(PORT))
}
