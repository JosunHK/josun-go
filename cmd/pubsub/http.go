package pubsub

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/JosunHK/josun-go.git/cmd/database"
	"github.com/JosunHK/josun-go.git/cmd/handlers/mahjong"
	mahjongStruct "github.com/JosunHK/josun-go.git/cmd/struct/mahjong"
	errorTemplate "github.com/JosunHK/josun-go.git/web/templates/contents/errorAlert"
	mahjongTemplates "github.com/JosunHK/josun-go.git/web/templates/contents/mahjong"
	watermillhttp "github.com/ThreeDotsLabs/watermill-http/v2/pkg/http"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	eventBus *cqrs.EventBus
}

func NewHandler(e *echo.Echo, eventBus *cqrs.EventBus, sseRouter watermillhttp.SSERouter) {
	h := Handler{
		eventBus: eventBus,
	}

	marshaler := cqrs.JSONMarshaler{}
	topic := marshaler.Name(mahjongStruct.GameStateUpdated{})
	stateHandler := sseRouter.AddHandler(topic, &Streamer{db: database.DB})

	e.POST("/mahjong/updateScore/:code", h.UpdateScore)
	e.GET("/mahjong/room/:code/state", func(c echo.Context) error {
		c.Request().SetPathValue("code", c.Param("code"))
		stateHandler(c.Response(), c.Request())
		return nil
	})
}

func (h Handler) UpdateScore(c echo.Context) error {
	roomCode := c.Param("code")
	if roomCode == "" {
		return errorTemplate.ErrorToast("Invalid Request").Render(c.Request().Context(), c.Response().Writer)
	}

	if err := mahjong.UpdateScore(c); err != nil {
		log.Error("UpdateScore error: " + err.Error())
		return errorTemplate.ErrorToast(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}

	event := mahjongStruct.GameUpdated{
		RoomCode: roomCode,
	}

	if err := h.eventBus.Publish(c.Request().Context(), event); err != nil {
		return errorTemplate.ErrorToast("Interal Server Error, try again later!").Render(c.Request().Context(), c.Response().Writer)
	}

	log.Info("UpdateScore for room : " + roomCode)
	return c.NoContent(http.StatusOK)
}

type Streamer struct {
	db *sql.DB
}

func (s *Streamer) InitialStreamResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	log.Info("sending InitStreamResponse")
	roomCode := r.PathValue("code")
	if roomCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid room ID"))
		return nil, false
	}

	resp, err := s.getResponse(r.Context(), roomCode, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return nil, false
	}

	return resp, true
}

func (s *Streamer) NextStreamResponse(r *http.Request, msg *message.Message) (response interface{}, ok bool) {
	log.Info("sending NextStreamResponse")
	roomCode := r.PathValue("code")
	if roomCode == "" {
		log.Error("invalid room ID")
		return nil, false
	}

	var event mahjongStruct.GameStateUpdated
	err := json.Unmarshal(msg.Payload, &event)
	if err != nil {
		log.Error("cannot unmarshal: " + err.Error())
		return "", false
	}

	if event.RoomCode != roomCode {
		return "", false
	}

	resp, err := s.getResponse(r.Context(), roomCode, &event)
	if err != nil {
		log.Error("could not get response: " + err.Error())
		return nil, false
	}

	return resp, true
}

func (s *Streamer) getResponse(ctx context.Context, roomCode string, event *mahjongStruct.GameStateUpdated) (interface{}, error) {
	var buffer bytes.Buffer
	var err error
	if event == nil {
		err = mahjongTemplates.InitRes().Render(ctx, &buffer)
	} else {
		err = mahjongTemplates.Update(*event).Render(ctx, &buffer)
	}

	if err != nil {
		return nil, err
	}

	return buffer.String(), nil
}
