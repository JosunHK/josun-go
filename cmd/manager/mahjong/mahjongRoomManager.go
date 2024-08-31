package mahjongManager

import (
	"fmt"
	"math/rand/v2"

	"github.com/JosunHK/josun-go.git/cmd/database"
	"github.com/JosunHK/josun-go.git/cmd/util/cookie"
	sqlc "github.com/JosunHK/josun-go.git/db/generated"
	"github.com/labstack/echo/v4"
)

const CODE_UPPER_BOUND = 9999
const CODE_LOWER_BOUND = 0

func CreateGameState(c echo.Context) (int64, error) {
	params := sqlc.CreateMahjongGameStateParams{
		RoundWind: sqlc.MahjongGameStateRoundWindEast,
		SeatWind:  sqlc.MahjongGameStateSeatWindEast,
		Round:     0,
	}

	DB := database.DB
	queries := sqlc.New(DB)

	result, err := queries.CreateMahjongGameState(c.Request().Context(), params)
	if err != nil {
		return 0, fmt.Errorf("Unable to create GameState", err)
	}

	return result.LastInsertId()
}

func CreateRoomOwner(c echo.Context) (int64, error) {
	roomOwnerParams, err := CreateMahjongRoomOwnerParams(c)
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

func CreateMahjongRoomOwnerParams(c echo.Context) (sqlc.CreateMahjongRoomOwnerParams, error) {
	guestID, err := cookie.GetGuestSessionUUID(c)
	if err != nil {
		return sqlc.CreateMahjongRoomOwnerParams{}, err
	}

	params := sqlc.CreateMahjongRoomOwnerParams{
		UserID:  0,
		GuestID: guestID,
	}

	return params, nil
}

func CreateMahjongPlayer(c echo.Context, params sqlc.CreateMahjongPlayerParams) (int64, error) {
	DB := database.DB
	queries := sqlc.New(DB)

	result, err := queries.CreateMahjongPlayer(c.Request().Context(), params)
	if err != nil {
		return 0, fmt.Errorf("Unable to create MahjongPlayer", err)
	}

	return result.LastInsertId()

}

func GetRandomRoomCode(c echo.Context) string {
	code := genCode()
	DB := database.DB
	queries := sqlc.New(DB)

	for _, err := queries.GetRoomWithCode(c.Request().Context(), code); err == nil; {
		code = genCode()
	}

	return code
}

func genCode() string {
	code := rand.IntN(CODE_UPPER_BOUND+1-CODE_LOWER_BOUND) + CODE_LOWER_BOUND
	return fmt.Sprintf("%04d", code)
}

func GetWindByIndex(index int) sqlc.MahjongPlayerWind {
	switch index {
	case 0:
		return sqlc.MahjongPlayerWindEast
	case 1:
		return sqlc.MahjongPlayerWindSouth
	case 2:
		return sqlc.MahjongPlayerWindWest
	case 3:
		return sqlc.MahjongPlayerWindNorth
	}

	return sqlc.MahjongPlayerWindEast

}
