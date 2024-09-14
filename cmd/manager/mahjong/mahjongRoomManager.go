package mahjongManager

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/JosunHK/josun-go.git/cmd/database"
	ms "github.com/JosunHK/josun-go.git/cmd/struct/mahjong"
	"github.com/JosunHK/josun-go.git/cmd/util/cookie"
	fp "github.com/JosunHK/josun-go.git/cmd/util/fp"
	sqlc "github.com/JosunHK/josun-go.git/db/generated"
	"github.com/labstack/echo/v4"
)

const CODE_UPPER_BOUND = 9999
const CODE_LOWER_BOUND = 0

var WIND_LIST = [4]string{
	string(sqlc.MahjongGameStateRoundWindEast),
	string(sqlc.MahjongGameStateRoundWindNorth),
	string(sqlc.MahjongGameStateRoundWindSouth),
	string(sqlc.MahjongGameStateRoundWindWest),
}

type Winner struct {
	fWinner ms.Winner
	rWinner *sqlc.MahjongPlayer
}

func CreateGameState(c echo.Context, queries *sqlc.Queries) (int64, error) {
	params := sqlc.CreateMahjongGameStateParams{
		RoundWind: sqlc.MahjongGameStateRoundWindEast,
		SeatWind:  sqlc.MahjongGameStateSeatWindEast,
		Round:     0,
	}

	if queries == nil {
		DB := database.DB
		queries = sqlc.New(DB)
	}

	result, err := queries.CreateMahjongGameState(c.Request().Context(), params)
	if err != nil {
		return 0, fmt.Errorf("Unable to create GameState", err)
	}

	return result.LastInsertId()
}

func GetOrCreateRoomOwner(c echo.Context, queries *sqlc.Queries) (int64, error) {
	roomOwnerParams, err := CreateMahjongRoomOwnerParams(c)
	if err != nil {
		return 0, fmt.Errorf("Unable to create params for room owner", err)
	}

	if queries == nil {
		queries = sqlc.New(database.DB)
	}

	user, err := queries.GetOwnerByUUIDorUserId(c.Request().Context(), sqlc.GetOwnerByUUIDorUserIdParams{
		GuestID: roomOwnerParams.GuestID,
		UserID:  roomOwnerParams.UserID,
	})
	if err == nil {
		return user.ID, nil
	}

	result, err := queries.CreateMahjongRoomOwner(c.Request().Context(), roomOwnerParams)
	if err != nil {
		return 0, fmt.Errorf("Unable to create RoomOwner", err)
	}

	return result.LastInsertId()
}

func GetGameData(c echo.Context, code string) (ms.GameData, error) {
	return GetGameDataWithContext(c.Request().Context(), code)
}

func GetGameDataWithContext(c context.Context, code string) (ms.GameData, error) {
	room, err := GetRoomByCode(c, code)
	if err != nil {
		return ms.GameData{}, err
	}

	state, err := GetGameStateByRoomCode(c, code)
	if err != nil {
		return ms.GameData{}, err
	}

	players, err := GetPlayersByRoomCode(c, code)
	if err != nil {
		return ms.GameData{}, err
	}

	return ms.GameData{
		Room:      room,
		GameState: state,
		Players:   players,
	}, nil
}

func CreateRoom(c echo.Context, queries *sqlc.Queries, roomParams sqlc.CreateMahjongRoomParams) (int64, error) {
	if queries == nil {
		queries = sqlc.New(database.DB)
	}

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

func CreateMahjongPlayer(c echo.Context, queries *sqlc.Queries, params sqlc.CreateMahjongPlayerParams) (int64, error) {
	if queries == nil {
		queries = sqlc.New(database.DB)
	}

	result, err := queries.CreateMahjongPlayer(c.Request().Context(), params)
	if err != nil {
		return 0, fmt.Errorf("Unable to create MahjongPlayer", err)
	}

	return result.LastInsertId()

}

func GetRandomRoomCode(c echo.Context) string {
	code := genCode()
	queries := sqlc.New(database.DB)

	for _, err := queries.GetRoomByCode(c.Request().Context(), code); err == nil; {
		code = genCode()
	}

	return code
}

func GetPlayersByRoomCode(c context.Context, code string) ([]sqlc.MahjongPlayer, error) {
	queries := sqlc.New(database.DB)

	players, err := queries.GetPlayersByRoomCode(c, code)
	if err != nil {
		return []sqlc.MahjongPlayer{}, fmt.Errorf("Unable to get game state with code", err)
	}

	return players, nil
}

func GetGameStateByRoomCode(c context.Context, code string) (sqlc.MahjongGameState, error) {
	queries := sqlc.New(database.DB)

	state, err := queries.GetGameStateByRoomCode(c, code)
	if err != nil {
		return sqlc.MahjongGameState{}, fmt.Errorf("Unable to get game state with code", err)
	}

	return state, nil
}

func GetRoomById(c echo.Context, Id int64) (sqlc.MahjongRoom, error) {
	queries := sqlc.New(database.DB)

	room, err := queries.GetRoomById(c.Request().Context(), Id)
	if err != nil {
		return sqlc.MahjongRoom{}, fmt.Errorf("Unable to get room with Id", err)
	}

	return room, nil
}

func GetRoomByCode(c context.Context, code string) (sqlc.MahjongRoom, error) {
	queries := sqlc.New(database.DB)

	room, err := queries.GetRoomByCode(c, code)
	if err != nil {
		return sqlc.MahjongRoom{}, fmt.Errorf("Unable to get room with code", err)
	}

	return room, nil
}

func GetPlayerById(c context.Context, id int64) (sqlc.MahjongPlayer, error) {
	queries := sqlc.New(database.DB)

	player, err := queries.GetPlayerById(c, id)
	if err != nil {
		return sqlc.MahjongPlayer{}, fmt.Errorf("Unable to get room with code", err)
	}

	return player, nil
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

func UpdatePlayerScore(c context.Context, queries *sqlc.Queries, id int64, score int32) error {
	if queries == nil {
		queries = sqlc.New(database.DB)
	}

	oriPlayer, err := queries.GetPlayerById(c, id)
	if err != nil {
		return fmt.Errorf("Player does not exist", err)
	}

	_, err = queries.CreateActionLog(c, sqlc.CreateActionLogParams{
		RoomID:     oriPlayer.RoomID,
		PlayerID:   id,
		ScoreDelta: score - oriPlayer.Score,
	})

	err = queries.UpdatePlayerScore(c, sqlc.UpdatePlayerScoreParams{
		ID:    id,
		Score: score,
	})

	if err != nil {
		return fmt.Errorf("Unable to update player score", err)
	}

	return nil
}

func UpdateGameState(c context.Context, queries *sqlc.Queries, gameState sqlc.MahjongGameState) error {
	if queries == nil {
		queries = sqlc.New(database.DB)
	}

	params := sqlc.UpdateGameStateParams{
		RoundWind: gameState.RoundWind,
		SeatWind:  gameState.SeatWind,
		Round:     gameState.Round,
		Kyoutaku:  gameState.Kyoutaku,
		Ended:     gameState.Ended,
		ID:        gameState.ID,
	}

	err := queries.UpdateGameState(c, params)
	if err != nil {
		return fmt.Errorf("Unable to update game state", err)
	}

	return nil
}

func HandleGameWin(c echo.Context, winForm ms.WinForm, gameData *ms.GameData) error {
	tx, err := database.DB.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	_ = sqlc.New(tx)

	if winForm.IsTsumo {
		err := handleTsumo(c, winForm, gameData)
		if err != nil {
			return err
		}
	} else {
		err := handleDirect(c, winForm, gameData)
		if err != nil {
			return err
		}
	}

	handlePlayerRiichi(winForm.RiichiPlayers, gameData)

	if err := handleKyoutaku(c, winForm, gameData); err != nil {
		if err != nil {
			return err
		}
	}

	for _, player := range gameData.Players {
		err = UpdatePlayerScore(c.Request().Context(), nil, player.ID, player.Score)
		if err != nil {
			return err
		}
	}

	err = UpdateGameState(c.Request().Context(), nil, gameData.GameState)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func HandleGameDraw(c echo.Context, drawForm ms.DrawForm, gameData *ms.GameData) error {
	tx, err := database.DB.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	queries := sqlc.New(tx)

	needAdvance, err := updateDrawGameStateReturnNeedAdvance(drawForm.DrawPlayers, drawForm.RiichiPlayers, gameData)
	if err != nil {
		return err
	}

	if needAdvance {
		advanceGameWind(gameData, scoreThreshold(gameData.Room.GameLength))
	} else {
		gameData.GameState.Round += 1
	}

	err = UpdateGameState(c.Request().Context(), queries, gameData.GameState)

	for _, player := range gameData.Players {
		err = UpdatePlayerScore(c.Request().Context(), queries, player.ID, player.Score)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func scoreThreshold(gameLength sqlc.MahjongRoomGameLength) int32 {
	if gameLength == sqlc.MahjongRoomGameLengthTonpuu {
		return 40000
	}
	return 30000
}

func handlePlayerRiichi(riichiPlayers []ms.RiichiPlayer, gameData *ms.GameData) {
	for i := range gameData.Players {
		player := &gameData.Players[i]
		for _, rplayer := range riichiPlayers {
			if player.ID == rplayer.PlayerId && rplayer.Riichi {
				player.Score -= int32(1000)
			}
		}
	}
}

func updateDrawGameStateReturnNeedAdvance(rawPlayers []ms.DrawPlayer, riichiPlayers []ms.RiichiPlayer, gameData *ms.GameData) (bool, error) {
	noTenCount := 0
	riichiCount := 0
	noTenMap := make(map[int64]bool)

	for _, p := range riichiPlayers {
		if p.Riichi {
			riichiCount++
		}
	}

	gameData.GameState.Kyoutaku += int32(riichiCount)
	handlePlayerRiichi(riichiPlayers, gameData)

	for _, rawPlayer := range rawPlayers {
		if rawPlayer.NoTen {
			for _, rplayer := range riichiPlayers {
				if rawPlayer.PlayerId == rplayer.PlayerId && rplayer.Riichi {
					return false, fmt.Errorf("Invalid draw")
				}
			}
			noTenCount++
		}

		noTenMap[rawPlayer.PlayerId] = rawPlayer.NoTen
	}

	switch noTenCount {
	case 1:
		for i := range gameData.Players {
			player := &gameData.Players[i]
			if !noTenMap[player.ID] {
				player.Score += 1000
			} else {
				player.Score -= 3000
			}
		}
	case 2:
		for i := range gameData.Players {
			player := &gameData.Players[i]
			if !noTenMap[player.ID] {
				player.Score += 1500
			} else {
				player.Score -= 1500
			}
		}
	case 3:
		for i := range gameData.Players {
			player := &gameData.Players[i]
			if !noTenMap[player.ID] {
				player.Score -= 3000
			} else {
				player.Score += 1000
			}
		}
	}

	for _, player := range gameData.Players {
		if noTenMap[player.ID] && (string(player.Wind) == string(gameData.GameState.SeatWind)) {
			return true, nil
		}
	}

	return false, nil
}

func genCode() string {
	code := rand.IntN(CODE_UPPER_BOUND+1-CODE_LOWER_BOUND) + CODE_LOWER_BOUND
	return fmt.Sprintf("%04d", code)
}

func advanceGameWind(gameData *ms.GameData, threshold int32) {
	gameState := &gameData.GameState
	players := gameData.Players
	gameLength := gameData.Room.GameLength

	newSeatWind, needUpdateRoundWind := getNextSeatWind(gameState.SeatWind)
	gameState.SeatWind = newSeatWind
	gameState.Round = 0

	if !needUpdateRoundWind {
		return
	}

	switch gameState.RoundWind {
	case sqlc.MahjongGameStateRoundWindEast:
		if gameLength == sqlc.MahjongRoomGameLengthTonpuu && isAboveEndThreshold(players, threshold) {
			gameState.Ended = true
		} else {
			gameState.RoundWind = sqlc.MahjongGameStateRoundWindSouth
		}
	case sqlc.MahjongGameStateRoundWindSouth:
		if gameLength == sqlc.MahjongRoomGameLengthTonpuu || isAboveEndThreshold(players, threshold) {
			gameState.Ended = true
		} else {
			gameState.RoundWind = sqlc.MahjongGameStateRoundWindWest
		}
	case sqlc.MahjongGameStateRoundWindWest:
		gameState.Ended = true
	}
}

func getNextSeatWind(seatWind sqlc.MahjongGameStateSeatWind) (newSeatWind sqlc.MahjongGameStateSeatWind, needUpdateRoundWind bool) {
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

func isAboveEndThreshold(players []sqlc.MahjongPlayer, threshold int32) bool {
	for _, player := range players {
		if player.Score >= threshold {
			return true
		}
	}
	return false
}

func mapToWinner(winForm ms.WinForm, gameData *ms.GameData) []Winner {
	winners := []Winner{}
	for _, fplayer := range winForm.Winners {
		for i, rplayer := range gameData.Players {
			if fplayer.PlayerId == rplayer.ID {
				winners = append(winners, Winner{fplayer, &gameData.Players[i]})
			}
		}
	}
	return winners
}

func handleDirect(c echo.Context, winForm ms.WinForm, gameData *ms.GameData) error {
	winners := mapToWinner(winForm, gameData)

	if len(winners) == 0 {
		return fmt.Errorf("Winner not found")
	}

	losers := fp.Filter2(func(p sqlc.MahjongPlayer) bool {
		return p.ID == winForm.LoserId
	}, &gameData.Players)

	if len(losers) == 0 || len(losers) > 1 {
		return fmt.Errorf("Invalid loser count")
	}

	loser := losers[0]
	needAdvanceGameWind := false

	for _, winner := range winners {
		scoreMap, err := getScoreMap(winner.fWinner.Han, winner.fWinner.Fu)
		if err != nil {
			return err
		}

		var score int
		if string(winner.rWinner.Wind) == string(gameData.GameState.SeatWind) {
			score = scoreMap.OyaDirect
		} else {
			score = scoreMap.KoDirect
		}

		if score == 0 {
			return fmt.Errorf("Invalid score")
		}

		winner.rWinner.Score += (int32(score) + gameData.GameState.Round*300)
		loser.Score -= (int32(score) + gameData.GameState.Round*300)

		needAdvanceGameWind = needAdvanceGameWind || string(winner.rWinner.Wind) == string(gameData.GameState.SeatWind)
	}

	if needAdvanceGameWind {
		gameData.GameState.Round += 1
	} else {
		advanceGameWind(gameData, scoreThreshold(gameData.Room.GameLength))
	}

	return nil
}

func handleTsumo(c echo.Context, winForm ms.WinForm, gameData *ms.GameData) error {
	var rWinner *sqlc.MahjongPlayer
	fWinner := winForm.Winners[0]

	for i, player := range gameData.Players {
		if player.ID == fWinner.PlayerId {
			rWinner = &gameData.Players[i]
			break
		}
	}

	if rWinner == nil {
		return fmt.Errorf("Winner not found")
	}

	scoreMap, err := getScoreMap(fWinner.Han, fWinner.Fu)
	if err != nil {
		return err
	}

	if string(rWinner.Wind) == string(gameData.GameState.SeatWind) {
		score := scoreMap.OyaTsumo
		if score == 0 {
			return fmt.Errorf("Invalid score")
		}

		rWinner.Score += int32(score*3) + gameData.GameState.Round*100*3
		for i, player := range gameData.Players {
			if player.ID != rWinner.ID {
				gameData.Players[i].Score -= (int32(score) + gameData.GameState.Round*100)
			}
		}
	} else {
		scoreKo := scoreMap.KoTsumoKo
		scoreOya := scoreMap.KoTsumoOya
		if scoreKo == 0 || scoreOya == 0 {
			return fmt.Errorf("Invalid score")
		}

		rWinner.Score += (int32(scoreKo)*2 + int32(scoreOya) + gameData.GameState.Round*100*3)
		for i, player := range gameData.Players {
			if player.ID == rWinner.ID {
				continue
			}
			if string(player.Wind) == string(gameData.GameState.SeatWind) {
				gameData.Players[i].Score -= (int32(scoreOya) + gameData.GameState.Round*100)
			} else {
				gameData.Players[i].Score -= (int32(scoreKo) + gameData.GameState.Round*100)
			}
		}
	}

	if string(rWinner.Wind) == string(gameData.GameState.SeatWind) {
		gameData.GameState.Round += 1
	} else {
		advanceGameWind(gameData, scoreThreshold(gameData.Room.GameLength))
	}

	return nil
}

func handleKyoutaku(c echo.Context, winForm ms.WinForm, gameData *ms.GameData) error {
	if winForm.IsTsumo {
		var kyoutakuWinner *sqlc.MahjongPlayer
		for i, player := range gameData.Players {
			if player.ID == winForm.Winners[0].PlayerId {
				kyoutakuWinner = &gameData.Players[i]
				break
			}
		}

		if kyoutakuWinner == nil {
			return fmt.Errorf("Kyoutaku winner not found")
		}

		kyoutakuWinner.Score += gameData.GameState.Kyoutaku * 1000
		gameData.GameState.Kyoutaku = 0
		return nil
	}

	var kyoutakuWinner *sqlc.MahjongPlayer
	winners := mapToWinner(winForm, gameData)

	losers := fp.Filter2(func(p sqlc.MahjongPlayer) bool {
		return p.ID == winForm.LoserId
	}, &gameData.Players)

	if len(losers) == 0 || len(losers) > 1 {
		return fmt.Errorf("Invalid loser count")
	}

	loser := losers[0]

	loserIndex := fp.IndexOf(string(loser.Wind), WIND_LIST[:])
	expectedKamicha := (loserIndex + 1) % 4

	for i := 0; i < 3; i++ {
		for _, winner := range winners {
			if fp.IndexOf(string(winner.rWinner.Wind), WIND_LIST[:]) == expectedKamicha {
				kyoutakuWinner = winner.rWinner
				break
			}
		}
		expectedKamicha = (expectedKamicha + 1) % 4
	}

	if kyoutakuWinner == nil {
		return fmt.Errorf("Kyoutaku winner not found")
	}

	kyoutakuWinner.Score += gameData.GameState.Kyoutaku * 1000
	gameData.GameState.Kyoutaku = 0

	return nil
}

func getScoreMap(han, fu int) (ms.Score, error) {
	if han == 1 && (fu == 20 || fu == 25) {
		return ms.Score{}, fmt.Errorf("Invalid han and fu")
	}

	return ms.ScoreMap[han][fu], nil
}

func GetInitGameState(c context.Context, code string) (ms.GameStateUpdated, error) {
	gameData, err := GetGameDataWithContext(c, code)
	if err != nil {
		return ms.GameStateUpdated{}, err
	}

	return ms.GameStateUpdated{
		RoomCode:  code,
		GameState: gameData.GameState,
		Players:   gameData.Players,
	}, nil
}
