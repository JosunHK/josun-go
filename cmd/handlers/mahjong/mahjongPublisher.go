package mahjong

import (
	"fmt"

	"github.com/JosunHK/josun-go.git/cmd/database"
	manager "github.com/JosunHK/josun-go.git/cmd/manager/mahjong"
	mahjongStruct "github.com/JosunHK/josun-go.git/cmd/struct/mahjong"
	"github.com/JosunHK/josun-go.git/cmd/util/cookie"
	sqlc "github.com/JosunHK/josun-go.git/db/generated"

	"github.com/labstack/echo/v4"
)

func UpdateScore(c echo.Context) error {
	code := c.Param("code")
	updateType := c.FormValue("updateType")

	gameData, err := validateRoomReturnGameData(c, code, updateType)
	if err != nil {
		return fmt.Errorf("Invalid updateType ", err)
	}

	switch updateType {
	case "win":
		err = UpdateScoreWin(c, &gameData)
	case "draw":
		err = UpdateScoreDraw(c, &gameData)
	case "manual":
		err = UpdateScoreManual(c, &gameData)
	}

	if err != nil {
		return fmt.Errorf("Invalid updateType ", err)
	}

	return nil
}

// prevent updating players in another room by modifiying the id in the request
func validatePlayers(c echo.Context, ids []int64, roomId int64) error {
	DB := database.DB
	queries := sqlc.New(DB)
	count, err := queries.GetPlayerCountByRoomId(c.Request().Context(), sqlc.GetPlayerCountByRoomIdParams{
		Ids:    ids,
		RoomID: roomId,
	})

	if err != nil {
		return fmt.Errorf("Unable to get player count", err)
	}

	if count != int64(len(ids)) {
		return fmt.Errorf("Invalid player id")
	}

	return nil
}

func validateRoomReturnGameData(c echo.Context, code, updateType string) (mahjongStruct.GameData, error) {
	DB := database.DB
	queries := sqlc.New(DB)

	if code == "" || updateType == "" {
		return mahjongStruct.GameData{}, fmt.Errorf("Invalid request")
	}

	gameData, err := manager.GetGameData(c, code)
	if err != nil {
		return mahjongStruct.GameData{}, fmt.Errorf("Room not found")
	}

	uuid, err := cookie.GetGuestSessionUUID(c)
	if err != nil {
		return mahjongStruct.GameData{}, fmt.Errorf("Invalid session")
	}

	owner, err := queries.GetOwnerByUUIDorUserId(c.Request().Context(), sqlc.GetOwnerByUUIDorUserIdParams{GuestID: uuid})
	if err != nil {
		return mahjongStruct.GameData{}, fmt.Errorf("Owner not found")
	}

	if gameData.Room.OwnerID != owner.ID {
		return mahjongStruct.GameData{}, fmt.Errorf("Unauthorized")
	}

	return gameData, nil
}

func UpdateScoreWin(c echo.Context, gameData *mahjongStruct.GameData) error {
	var winForm mahjongStruct.WinForm
	err := decoder.Decode(&winForm, c.Request().PostForm)
	if err != nil {
		return fmt.Errorf("Failed to decode drawform", err)
	}

	ids := []int64{winForm.WinnerId}
	if !winForm.IsTsumo {
		ids = append(ids, winForm.LoserId)
	}

	if err := validatePlayers(c, ids, gameData.Room.ID); err != nil {
		err = fmt.Errorf("Invalid player", err)
		return err
	}

	if err := manager.HandleGameWin(c, winForm, gameData); err != nil {
		return fmt.Errorf("Failed to handle game draw", err)
	}

	return nil
}

func UpdateScoreDraw(c echo.Context, gameData *mahjongStruct.GameData) error {
	var drawForm mahjongStruct.DrawForm
	err := decoder.Decode(&drawForm, c.Request().PostForm)
	if err != nil {
		return fmt.Errorf("Failed to decode drawform", err)
	}

	ids := []int64{}
	for _, player := range drawForm.DrawPlayers {
		ids = append(ids, player.PlayerId)
	}

	if err := validatePlayers(c, ids, gameData.Room.ID); err != nil {
		err = fmt.Errorf("Invalid player", err)
		return err
	}

	if err := manager.HandleGameDraw(c, drawForm, gameData); err != nil {
		return fmt.Errorf("Failed to handle game draw", err)
	}

	return nil
}

func UpdateScoreManual(c echo.Context, gameData *mahjongStruct.GameData) error {
	DB := database.DB
	queries := sqlc.New(DB)

	var manualForm mahjongStruct.ManualForm
	if err := decoder.Decode(&manualForm, c.Request().PostForm); err != nil {
		return fmt.Errorf("Failed to decode roomSetting", err)
	}

	if err := validatePlayers(c, []int64{manualForm.PlayerId}, gameData.Room.ID); err != nil {
		err = fmt.Errorf("Invalid player", err)
		return err
	}

	player := sqlc.MahjongPlayer{}
	for _, p := range gameData.Players {
		if p.ID == manualForm.PlayerId {
			player = p
			break
		}
	}

	if player.ID == 0 {
		return fmt.Errorf("Failed to get Player")
	}

	err := queries.UpdatePlayerScore(c.Request().Context(), sqlc.UpdatePlayerScoreParams{
		Score: player.Score + int32(manualForm.Score),
		ID:    player.ID,
	})
	if err != nil {
		return fmt.Errorf("Failed to update Player", err)
	}

	return nil
}