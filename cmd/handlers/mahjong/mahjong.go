package mahjong

import (
	"fmt"
	"math/rand"

	"github.com/JosunHK/josun-go.git/cmd/database"
	mgr "github.com/JosunHK/josun-go.git/cmd/manager/mahjong"
	responseUtil "github.com/JosunHK/josun-go.git/cmd/util/response"
	sqlc "github.com/JosunHK/josun-go.git/db/generated"
	errorTemplate "github.com/JosunHK/josun-go.git/web/templates/contents/errorAlert"
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

func Room(c echo.Context) templ.Component {
	code := c.Param("code")
	DB := database.DB
	queries := sqlc.New(DB)

	players, err := queries.GetPlayersByRoomCode(c.Request().Context(), code)
	if err != nil || len(players) == 0 {
		return errorTemplate.ErrorAlert("Room Not Found", "The room you are looking for does not exist.")
	}

	gameState, err := queries.GetGameStateByRoomCode(c.Request().Context(), code)
	if err != nil || len(players) == 0 {
		return errorTemplate.ErrorAlert("Room Not Found", "The room you are looking for does not exist.")
	}

	return mahjongTemplates.Room(players, code, gameState)
}

func RoomCreate(c echo.Context) (string, error) {
	DB := database.DB
	tx, err := DB.Begin()
	if err != nil {
		err = fmt.Errorf("Failed to begin transaction", err)
		log.Error(err)
		return "", err
	}

	defer tx.Rollback()

	queries := sqlc.New(tx)

	type RoomSetting struct {
		PlayerNames []string                   `schema:"playerNames"`
		GameLength  sqlc.MahjongRoomGameLength `schema:"gameLength,required"`
		StartPoints int                        `schema:"startPoints,required"`
	}

	if err := c.Request().ParseForm(); err != nil {
		return "", fmt.Errorf("Failed to parse form", err)
	}

	var roomSetting RoomSetting

	err = decoder.Decode(&roomSetting, c.Request().PostForm)
	if err != nil {
		return "", fmt.Errorf("Failed to decode roomSetting", err)
	}

	if roomSetting.GameLength == sqlc.MahjongRoomGameLengthHanChan {
		fillOutWithRandomNames(&roomSetting.PlayerNames, 4)
	}

	if roomSetting.GameLength == sqlc.MahjongRoomGameLengthTonpuu {
		fillOutWithRandomNames(&roomSetting.PlayerNames, 3)
	}

	ownerId, err := mgr.GetOrCreateRoomOwner(c, queries)
	if err != nil {
		return "", fmt.Errorf("Failed to create room owner", err)
	}

	stateId, err := mgr.CreateGameState(c, queries)
	if err != nil {
		return "", fmt.Errorf("Failed to create game state", err)
	}

	roomCode := mgr.GetRandomRoomCode(c)
	roomParams := sqlc.CreateMahjongRoomParams{
		GameStateID: stateId,
		RoomCode:    roomCode,
		GameLength:  roomSetting.GameLength,
		OwnerID:     ownerId,
	}

	roomId, err := mgr.CreateRoom(c, queries, roomParams)
	if err != nil {
		return "", fmt.Errorf("Failed to create room", err)
	}

	for i, name := range roomSetting.PlayerNames {
		_, err := mgr.CreateMahjongPlayer(c, queries, sqlc.CreateMahjongPlayerParams{
			RoomID: roomId,
			Name:   name,
			Score:  int32(roomSetting.StartPoints),
			Wind:   mgr.GetWindByIndex(i),
		})

		if err != nil {
			return "", fmt.Errorf("Failed to create player", err)
		}
	}

	return fmt.Sprintf("/mahjong/room/%s", roomCode), tx.Commit()
}

func fillOutWithRandomNames(names *[]string, count int) {
	randomNames := []string{
		"Carriage Lau", "Dragon Slayer", "土田浩翔", "伊藤誠", "岩倉玲音",
		"小島秀夫", "耶穌", "奈須蘑菇", "Ryan Gosling", "藤丸立香", "牧瀬紅莉栖(AI)",
		"Stocking(1/999)", "成神陽太", "宮永咲", "小泉純一郎", "何屋未来", "野口英世",
		"Boris Johnson", "Xi Jinping", "Mao Zedong", "太空希特勒", "花京院 🍩", "一姫",
	}

	for i := len(*names); i < count; i++ {
		randomName := randomNames[rand.Intn(len(randomNames))]
		*names = append(*names, randomName)
	}
}
