package mahjongManager

import (
	"fmt"
	"math/rand/v2"

	"github.com/JosunHK/josun-go.git/cmd/database"
	mahjongStruct "github.com/JosunHK/josun-go.git/cmd/struct/mahjong"
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

func GetGameData(c echo.Context, code string) (mahjongStruct.GameData, error) {
	room, err := GetRoomByCode(c, code)
	if err != nil {
		return mahjongStruct.GameData{}, err
	}

	state, err := GetGameStateByRoomCode(c, code)
	if err != nil {
		return mahjongStruct.GameData{}, err
	}

	players, err := GetPlayersByRoomCode(c, code)
	if err != nil {
		return mahjongStruct.GameData{}, err
	}

	return mahjongStruct.GameData{
		Room:      room,
		GameState: state,
		Players:   players,
	}, nil
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

	for _, err := queries.GetRoomByCode(c.Request().Context(), code); err == nil; {
		code = genCode()
	}

	return code
}

func GetPlayersByRoomCode(c echo.Context, code string) ([]sqlc.MahjongPlayer, error) {
	DB := database.DB
	queries := sqlc.New(DB)

	players, err := queries.GetPlayersByRoomCode(c.Request().Context(), code)
	if err != nil {
		return []sqlc.MahjongPlayer{}, fmt.Errorf("Unable to get game state with code", err)
	}

	return players, nil
}

func GetGameStateByRoomCode(c echo.Context, code string) (sqlc.MahjongGameState, error) {
	DB := database.DB
	queries := sqlc.New(DB)

	state, err := queries.GetGameStateByRoomCode(c.Request().Context(), code)
	if err != nil {
		return sqlc.MahjongGameState{}, fmt.Errorf("Unable to get game state with code", err)
	}

	return state, nil
}

func GetRoomById(c echo.Context, Id int64) (sqlc.MahjongRoom, error) {
	DB := database.DB
	queries := sqlc.New(DB)

	room, err := queries.GetRoomById(c.Request().Context(), Id)
	if err != nil {
		return sqlc.MahjongRoom{}, fmt.Errorf("Unable to get room with Id", err)
	}

	return room, nil
}

func GetRoomByCode(c echo.Context, code string) (sqlc.MahjongRoom, error) {
	DB := database.DB
	queries := sqlc.New(DB)

	room, err := queries.GetRoomByCode(c.Request().Context(), code)
	if err != nil {
		return sqlc.MahjongRoom{}, fmt.Errorf("Unable to get room with code", err)
	}

	return room, nil
}

func GetPlayerById(c echo.Context, id int64) (sqlc.MahjongPlayer, error) {
	DB := database.DB
	queries := sqlc.New(DB)

	player, err := queries.GetPlayerById(c.Request().Context(), id)
	if err != nil {
		return sqlc.MahjongPlayer{}, fmt.Errorf("Unable to get room with code", err)
	}

	return player, nil
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
func UpdatePlayerScore(c echo.Context, id int64, score int32) error {
	DB := database.DB
	queries := sqlc.New(DB)

	err := queries.UpdatePlayerScore(c.Request().Context(), sqlc.UpdatePlayerScoreParams{
		ID:    id,
		Score: score,
	})

	if err != nil {
		return fmt.Errorf("Unable to update player score", err)
	}

	return nil
}

func UpdateGameState(c echo.Context, gameState sqlc.MahjongGameState) error {
	DB := database.DB
	queries := sqlc.New(DB)

	params := sqlc.UpdateGameStateParams{
		RoundWind: gameState.RoundWind,
		SeatWind:  gameState.SeatWind,
		Round:     gameState.Round,
		Kyoutaku:  gameState.Kyoutaku,
		Ended:     gameState.Ended,
		ID:        gameState.ID,
	}

	err := queries.UpdateGameState(c.Request().Context(), params)
	if err != nil {
		return fmt.Errorf("Unable to update game state", err)
	}

	return nil
}

func AdvanceGameRound(c echo.Context, gameState sqlc.MahjongGameState) error {
	gameState.Round++
	err := UpdateGameState(c, gameState)
	if err != nil {
		return err
	}
	return nil
}

func GetNextSeatWind(seatWind sqlc.MahjongGameStateSeatWind) (newSeatWind sqlc.MahjongGameStateSeatWind, needUpdateRoundWind bool) {
	switch seatWind {
	case sqlc.MahjongGameStateSeatWindEast:
		return sqlc.MahjongGameStateSeatWindSouth, false
	case sqlc.MahjongGameStateSeatWindSouth:
		return sqlc.MahjongGameStateSeatWindWest, false
	case sqlc.MahjongGameStateSeatWindWest:
		return sqlc.MahjongGameStateSeatWindNorth, false
	case sqlc.MahjongGameStateSeatWindNorth:
		return sqlc.MahjongGameStateSeatWindEast, true
	}
	return sqlc.MahjongGameStateSeatWindEast, false
}

func IsAboveEndThreshold(players []sqlc.MahjongPlayer, threshold int32) bool {
	for _, player := range players {
		if player.Score >= threshold {
			return true
		}
	}
	return false
}

func AdvanceGameWind(c echo.Context, gameData *mahjongStruct.GameData, threshold int32) error {
	gameState := gameData.GameState
	players := gameData.Players
	gameLength := gameData.Room.GameLength

	newSeatWind, needUpdateRoundWind := GetNextSeatWind(gameState.SeatWind)
	gameState.SeatWind = newSeatWind
	if needUpdateRoundWind {
		gameState.Round = 0
		switch gameState.RoundWind {
		case sqlc.MahjongGameStateRoundWindEast:
			if gameLength == sqlc.MahjongRoomGameLengthTonpuu && IsAboveEndThreshold(players, threshold) {
				gameState.Ended = true
			} else {
				gameState.RoundWind = sqlc.MahjongGameStateRoundWindSouth
			}
		case sqlc.MahjongGameStateRoundWindSouth:
			if gameLength == sqlc.MahjongRoomGameLengthTonpuu || IsAboveEndThreshold(players, threshold) {
				gameState.Ended = true
			} else {
				gameState.RoundWind = sqlc.MahjongGameStateRoundWindWest
			}
		case sqlc.MahjongGameStateRoundWindWest:
			gameState.Ended = true
		}
	}

	if err := UpdateGameState(c, gameState); err != nil {
		return err
	}
	return nil
}

// I dont like this func, It would look so much better if it were written in Haskell
func HandleGameDraw(c echo.Context, rawPlayers []mahjongStruct.DrawPlayer, code string) error {
	noTenpaiCount := 0
	noTenpaiMap := make(map[int64]bool)

	gameData, err := GetGameData(c, code)
	if err != nil {
		return err
	}

	for _, rawPlayer := range rawPlayers {
		if !rawPlayer.Tenpai {
			noTenpaiCount++
		}

		noTenpaiMap[rawPlayer.PlayerId] = rawPlayer.Tenpai
	}

	switch noTenpaiCount {
	case 1:
		for _, player := range gameData.Players {
			if !noTenpaiMap[player.ID] {
				continue
			}
			err := UpdatePlayerScore(c, player.ID, player.Score-1000)
			if err != nil {
				return err
			}
		}
	case 2:
		for _, player := range gameData.Players {
			if !noTenpaiMap[player.ID] {
				continue
			}
			err := UpdatePlayerScore(c, player.ID, player.Score-1500)
			if err != nil {
				return err
			}
		}
	case 3:
		for _, player := range gameData.Players {
			if !noTenpaiMap[player.ID] {
				err := UpdatePlayerScore(c, player.ID, player.Score-3000)
				if err != nil {
					return err
				}
			}
		}
	}

	for _, player := range gameData.Players {
		if !noTenpaiMap[player.ID] && (string(player.Wind) == string(gameData.GameState.SeatWind)) {
			err := AdvanceGameRound(c, gameData.GameState)
			if err != nil {
				return err
			}
			return nil
		}
	}

	var threshold int32

	if gameData.Room.GameLength == sqlc.MahjongRoomGameLengthTonpuu {
		threshold = 40000
	} else {
		threshold = 30000
	}

	return AdvanceGameWind(c, &gameData, threshold)
}
