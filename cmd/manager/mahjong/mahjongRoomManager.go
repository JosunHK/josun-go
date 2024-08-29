package mahjongManager

import (
	"fmt"

	"github.com/JosunHK/josun-go.git/cmd/database"
	util "github.com/JosunHK/josun-go.git/cmd/util/mahjong"
	sqlc "github.com/JosunHK/josun-go.git/db/generated"
	"github.com/labstack/echo/v4"
)

func CreateRoomOwner(c echo.Context) (int64, error) {
	roomOwnerParams, err := util.CreateMahjongRoomOwnerParams(c)
	if err != nil {
		return 0, fmt.Errorf("Unable to create params for room owner", err)
	}

	DB := database.DB
	queries := sqlc.New(DB)

	result, err := queries.CreateMahjongRoomOwner(c.Request().Context(), roomOwnerParams)
	if err != nil {
		return 0, fmt.Errorf("Unable to create RoomOwner", err)
	}

	return result.LastInsertId()
}

func CreateRoom(c echo.Context, roomParams sqlc.CreateMahjongRoomParams) (int64, error) {
	DB := database.DB
	queries := sqlc.New(DB)

	result, err := queries.CreateMahjongRoom(c.Request().Context(), roomParams)
	if err != nil {
		return 0, fmt.Errorf("Unable to create Room", err)
	}

	return result.LastInsertId()
}
