package mahjong

import (
	"fmt"

	"github.com/JosunHK/josun-go.git/cmd/database"
	"github.com/JosunHK/josun-go.git/cmd/util/cookie"
	sqlc "github.com/JosunHK/josun-go.git/db/generated"

	"github.com/labstack/echo/v4"
)

func UpdateScore(c echo.Context) error {
	code := c.Param("code")
	updateType := c.FormValue("updateType")

	roomId, err := validateRoomReturnId(c, code, updateType)
	if err != nil {
		return fmt.Errorf("Invalid updateType ", err)
	}

	switch updateType {
	case "win":
		err = UpdateScoreWin(c, code, roomId)
	case "draw":
		err = UpdateScoreDraw(c, code, roomId)
	case "manual":
		err = UpdateScoreManual(c, code, roomId)
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

func validateRoomReturnId(c echo.Context, code, updateType string) (int64, error) {
	DB := database.DB
	queries := sqlc.New(DB)

	if code == "" || updateType == "" {
		return 0, fmt.Errorf("Invalid request")
	}

	room, err := queries.GetRoomByCode(c.Request().Context(), code)
	if err != nil {
		return 0, fmt.Errorf("Room not found")
	}

	uuid, err := cookie.GetGuestSessionUUID(c)
	if err != nil {
		return 0, fmt.Errorf("Invalid session")
	}

	owner, err := queries.GetOwnerByUUIDorUserId(c.Request().Context(), sqlc.GetOwnerByUUIDorUserIdParams{GuestID: uuid})
	if err != nil {
		return 0, fmt.Errorf("Owner not found")
	}

	if room.OwnerID != owner.ID {
		return 0, fmt.Errorf("Unauthorized")
	}

	return room.ID, nil
}

func UpdateScoreWin(c echo.Context, code string, roomId int64) error {
	type WinForm struct {
	}

	return nil
}

func UpdateScoreDraw(c echo.Context, code string, roomId int64) error {
	type DrawForm struct {
	}

	return nil
}

func UpdateScoreManual(c echo.Context, code string, roomId int64) error {
	DB := database.DB
	queries := sqlc.New(DB)

	type ManualForm struct {
		UpdateType string `schema:"updateType,required"`
		PlayerId   int64  `schema:"playerId,required"`
		Score      int    `schema:"score,default:0"`
	}

	var manualForm ManualForm
	err := decoder.Decode(&manualForm, c.Request().PostForm)
	if err != nil {
		return fmt.Errorf("Failed to decode roomSetting", err)
	}

	if err := validatePlayers(c, []int64{manualForm.PlayerId}, roomId); err != nil {
		err = fmt.Errorf("Invalid player", err)
		return err
	}

	player, err := queries.GetPlayerById(c.Request().Context(), manualForm.PlayerId)
	if err != nil {
		return fmt.Errorf("Failed to get Player", err)
	}

	player.Score += int32(manualForm.Score)

	err = queries.UpdatePlayerScore(c.Request().Context(), sqlc.UpdatePlayerScoreParams{
		Score: player.Score,
		ID:    player.ID,
	})
	if err != nil {
		return fmt.Errorf("Failed to update Player", err)
	}

	return nil
}
