package cookie

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const GUEST_SESSION_ID = "guestSessionId"

func ManageGuestSession(c *echo.Context) {
	cookie, err := (*c).Cookie(GUEST_SESSION_ID)
	if err != nil {
		CreateGuestSession(c)
	} else {
		RefreshGuestSession(cookie)
	}
}

func CreateGuestSession(c *echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = GUEST_SESSION_ID
	cookie.Value = uuid.New().String()
	cookie.Expires = time.Now().Add(240 * time.Hour) //10 days
	(*c).SetCookie(cookie)
}

func GetGuestSessionUUID(c echo.Context) (string, error) {
	cookie, err := c.Cookie(GUEST_SESSION_ID)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func RefreshGuestSession(cookie *http.Cookie) {
	cookie.Expires = time.Now().Add(240 * time.Hour)
}
