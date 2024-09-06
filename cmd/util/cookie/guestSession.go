package cookie

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
	cookie.Path = "/"
	cookie.Value = uuid.New().String()
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Secure = true
	cookie.Expires = time.Now().Add(240 * time.Hour) //10 days
	(*c).SetCookie(cookie)
}

func GetGuestSessionUUID(c echo.Context) (string, error) {
	cookie, err := c.Cookie(GUEST_SESSION_ID)
	if err != nil {
		log.Error("Error getting cookie: ", err)
		return "", err
	}

	return cookie.Value, nil
}

func RefreshGuestSession(cookie *http.Cookie) {
	cookie.Expires = time.Now().Add(240 * time.Hour)
}
