package mahjong

import (
	"fmt"

	mgr "github.com/JosunHK/josun-go.git/cmd/manager/mahjong"
	responseUtil "github.com/JosunHK/josun-go.git/cmd/util/response"
	sqlc "github.com/JosunHK/josun-go.git/db/generated"
	mahjongTemplates "github.com/JosunHK/josun-go.git/web/templates/contents/mahjong"
	"github.com/a-h/templ"
	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

var decoder = schema.NewDecoder()

func RoomSelect(c echo.Context) templ.Component {
	return mahjongTemplates.RoomSelect()
}

func RoomSetting(c echo.Context) error {
	return responseUtil.HTML(c, mahjongTemplates.RoomCreate())
}

func Test(c echo.Context) templ.Component {
	return mahjongTemplates.Test()
}

func RoomCreate(c echo.Context) error {
	type RoomSetting struct {
		PlayerNames []string `schema:"playerNames,default:P1|P2|P3|P4"`
		GameLength  string   `schema:"gameLength,required"`
		StartPoints int      `schema:"startPoints,required"`
	}

	err := c.Request().ParseForm()
	if err != nil {
		log.Error(fmt.Errorf("Failed to parse form", err))
		return err
	}

	var roomSetting RoomSetting

	err = decoder.Decode(&roomSetting, c.Request().PostForm)
	if err != nil {
		log.Error(fmt.Errorf("Failed to decode roomSetting", err))
		return err
	}

	ownerId, err := mgr.CreateRoomOwner(c)
	if err != nil {
		log.Error(fmt.Errorf("Failed to create room owner", err))
		return err
	}

	//TODO: create game state first
	//TODO: create craete generator for room code
	mgr.CreateRoom(c, sqlc.CreateMahjongRoomParams{
		GameStateID: 0,
		RoomCode:    "1234",
		GameLength:  roomSetting.GameLength,
		OwnerID:     ownerId,
	})

	return fmt.Errorf("Not implemented") // TODO: Implement this;
}
