package pubsub

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JosunHK/josun-go.git/cmd/database"
	dummyTemplates "github.com/JosunHK/josun-go.git/web/templates/contents/dummy"
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
	topic := marshaler.Name(GameStateUpdated{})
	stateHandler := sseRouter.AddHandler(topic, &Streamer{db: database.DB})

	e.POST("/mahjong/room/:code/state", h.UpdateScore)
	e.GET("/mahjong/room/:code/state", func(c echo.Context) error {
		c.Request().SetPathValue("code", c.Param("code"))
		stateHandler(c.Response(), c.Request())
		return nil
	})
}

func (h Handler) UpdateScore(c echo.Context) error {
	roomCode := c.Param("code")
	if roomCode == "" {
		return c.String(http.StatusBadRequest, "invalid room ID")
	}

	event := GameUpdated{
		RoomCode: roomCode,
	}

	err := h.eventBus.Publish(c.Request().Context(), event)
	if err != nil {
		return err
	}

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
		fmt.Println("invalid room ID")
		return nil, false
	}

	var event GameStateUpdated
	err := json.Unmarshal(msg.Payload, &event)
	if err != nil {
		fmt.Println("cannot unmarshal: " + err.Error())
		return "", false
	}

	if event.RoomCode != roomCode {
		return "", false
	}

	resp, err := s.getResponse(r.Context(), roomCode, &event)
	if err != nil {
		fmt.Println("could not get response: " + err.Error())
		return nil, false
	}

	return resp, true
}

func (s *Streamer) getResponse(ctx context.Context, roomCode string, event *GameStateUpdated) (interface{}, error) {
	var buffer bytes.Buffer
	var err error
	if event == nil {
		err = dummyTemplates.InitialRes().Render(ctx, &buffer)
	} else {
		err = dummyTemplates.Update(event.RoomCode).Render(ctx, &buffer)
	}

	if err != nil {
		return nil, err
	}

	return buffer.String(), nil
}
