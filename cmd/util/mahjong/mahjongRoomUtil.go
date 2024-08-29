package mahjongUtil

import (
	"github.com/JosunHK/josun-go.git/cmd/util/cookie"
	sqlc "github.com/JosunHK/josun-go.git/db/generated"
	"github.com/labstack/echo/v4"
)

// TODO: also allow create for auth user
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
