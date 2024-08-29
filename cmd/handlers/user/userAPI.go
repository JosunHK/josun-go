package user

import (
	"net/http"

	"github.com/JosunHK/josun-go.git/cmd/database"
	sqlc "github.com/JosunHK/josun-go.git/db/generated"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

func PostUser(c echo.Context) (err error, statusCode int, resObj interface{}) {
	type res struct {
		ID int64 `json:"id"`
	}

	DB := database.DB
	queries := sqlc.New(DB)

	result, err := queries.CreateUser(c.Request().Context(), sqlc.CreateUserParams{
		Name:     "test",
		Email:    "bruh@bruh",
		Password: "test",
	})

	if err != nil {
		return err, http.StatusInternalServerError, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err, http.StatusInternalServerError, nil
	}

	resObj = res{
		ID: id,
	}

	log.Info("User created with id: ", id)

	return nil, http.StatusCreated, resObj
}

func GetUsers(c echo.Context) (err error, statusCode int, resObj interface{}) {
	type response struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	DB := database.DB
	queries := sqlc.New(DB)

	users, err := queries.ListUsers(c.Request().Context())
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
