package api

import (
	"net/http"

	"github.com/JosunHK/josun-go.git/cmd/database"
	"github.com/JosunHK/josun-go.git/test"
	"github.com/labstack/echo/v4"
)

func GetUsers(ctx echo.Context) (err error, statusCode int, resObj interface{}) {
	type response struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	DB := database.DB
	queries := test.New(DB)

	users, err := queries.ListUsers(ctx.Request().Context())
	if err != nil {
		return err, http.StatusInternalServerError, nil
	}

	if len(users) == 0 {
		return nil, http.StatusNotFound, nil
	}

	resObj = response{
		ID:   users[0].ID,
		Name: users[0].Name,
	}

	return nil, http.StatusOK, resObj
}
