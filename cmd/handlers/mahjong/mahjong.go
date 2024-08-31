package mahjong

import (
	"fmt"

	"github.com/JosunHK/josun-go.git/cmd/database"
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

func RoomCreate(c echo.Context) (string, error) {
	DB := database.DB
	tx, err := DB.Begin()
	defer tx.Rollback()

	type RoomSetting struct {
		PlayerNames []string                   `schema:"playerNames,default:P1|P2|P3|P4"`
		GameLength  sqlc.MahjongRoomGameLength `schema:"gameLength,required"`
		StartPoints int                        `schema:"startPoints,required"`
	}

	if err := c.Request().ParseForm(); err != nil {
		err = fmt.Errorf("Failed to parse form", err)
		log.Error(err)
		return "", err
	}

	var roomSetting RoomSetting

	err = decoder.Decode(&roomSetting, c.Request().PostForm)
	if err != nil {
		err = fmt.Errorf("Failed to decode roomSetting", err)
		log.Error(err)
		return "", err
	}

	ownerId, err := mgr.CreateRoomOwner(c)
	if err != nil {
		err = fmt.Errorf("Failed to create room owner", err)
		log.Error(err)
		return "", err
	}

	stateId, err := mgr.CreateGameState(c)
	if err != nil {
		err = fmt.Errorf("Failed to create game state", err)
		log.Error(err)
		return "", err
	}

	roomCode := mgr.GetRandomRoomCode(c)
	roomParams := sqlc.CreateMahjongRoomParams{
		GameStateID: stateId,
		RoomCode:    roomCode,
		GameLength:  roomSetting.GameLength,
		OwnerID:     ownerId,
	}

	roomId, err := mgr.CreateRoom(c, roomParams)
	if err != nil {
		err = fmt.Errorf("Failed to create room", err)
		log.Error(err)
		return "", err
	}

	for i, name := range roomSetting.PlayerNames {
		_, err := mgr.CreateMahjongPlayer(c, sqlc.CreateMahjongPlayerParams{
			RoomID: roomId,
			Name:   name,
			Score:  int32(roomSetting.StartPoints),
			Wind:   mgr.GetWindByIndex(i),
		})

		if err != nil {
			err = fmt.Errorf("Failed to create player", err)
			log.Error(err)
			return "", err
		}
	}

	return fmt.Sprintf("/mahjong/room/%d", roomCode), nil
}
