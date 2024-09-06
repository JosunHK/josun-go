package pubsub

import (
	"context"
	"fmt"

	"github.com/JosunHK/josun-go.git/cmd/database"
	mahjongStruct "github.com/JosunHK/josun-go.git/cmd/struct/mahjong"
	sqlc "github.com/JosunHK/josun-go.git/db/generated"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-http/v2/pkg/http"
	"github.com/ThreeDotsLabs/watermill-sql/v3/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	log "github.com/sirupsen/logrus"
)

type Routers struct {
	EventsRouter *message.Router
	SSERouter    http.SSERouter
	EventBus     *cqrs.EventBus
}

func StartEventsRouter(ctx context.Context, routers Routers) {
	err := routers.EventsRouter.Run(context.Background())
	if err != nil {
		panic(err)
	}
}

func StartSSERouter(ctx context.Context, routers Routers) {
	err := routers.SSERouter.Run(context.Background())
	if err != nil {
		panic(err)
	}
}

func NewRouters() (Routers, error) {
	logger := watermill.NewStdLoggerWithOut(log.StandardLogger().Out, false, false)

	publisher, err := sql.NewPublisher(
		database.DB,
		sql.PublisherConfig{
			SchemaAdapter: sql.DefaultMySQLSchema{},
		},
		logger,
	)
	if err != nil {
		return Routers{}, err
	}

	//for publishing
	eventBus, err := cqrs.NewEventBusWithConfig(
		publisher,
		cqrs.EventBusConfig{
			GeneratePublishTopic: func(params cqrs.GenerateEventPublishTopicParams) (string, error) {
				return params.EventName, nil
			},
			Marshaler: cqrs.JSONMarshaler{},
			Logger:    logger,
		},
	)
	if err != nil {
		return Routers{}, err
	}

	eventsRouter, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return Routers{}, err
	}

	eventsRouter.AddMiddleware(middleware.Recoverer)

	//for subscribing
	eventProcessor, err := cqrs.NewEventProcessorWithConfig(
		eventsRouter,
		cqrs.EventProcessorConfig{
			GenerateSubscribeTopic: func(params cqrs.EventProcessorGenerateSubscribeTopicParams) (string, error) {
				return params.EventName, nil
			},
			SubscriberConstructor: func(params cqrs.EventProcessorSubscriberConstructorParams) (message.Subscriber, error) {
				return sql.NewSubscriber(
					database.DB,
					sql.SubscriberConfig{
						SchemaAdapter:    sql.DefaultMySQLSchema{},
						OffsetsAdapter:   sql.DefaultMySQLOffsetsAdapter{},
						InitializeSchema: true,
					},
					logger,
				)
			},
			Marshaler: cqrs.JSONMarshaler{},
			Logger:    logger,
		},
	)
	if err != nil {
		return Routers{}, err
	}

	err = eventProcessor.AddHandlers(
		cqrs.NewEventHandler(
			"UpdateGame",
			func(ctx context.Context, event *mahjongStruct.GameUpdated) error {
				gameState, err := GetGameStateEventByCode(ctx, event.RoomCode)
				if err != nil {
					log.Info("Error getting game state event by code\n", err)
					return err
				}
				return eventBus.Publish(ctx, gameState)
			},
		),
	)

	if err != nil {
		return Routers{}, err
	}

	sseSubscriber, err := sql.NewSubscriber(
		database.DB,
		sql.SubscriberConfig{
			SchemaAdapter:    sql.DefaultMySQLSchema{},
			OffsetsAdapter:   sql.DefaultMySQLOffsetsAdapter{},
			InitializeSchema: true,
		},
		logger,
	)
	if err != nil {
		return Routers{}, err
	}

	sseRouter, err := http.NewSSERouter(
		http.SSERouterConfig{
			UpstreamSubscriber: sseSubscriber,
			Marshaler:          http.StringSSEMarshaler{},
		},
		logger,
	)
	if err != nil {
		return Routers{}, err
	}

	return Routers{
		EventsRouter: eventsRouter,
		SSERouter:    sseRouter,
		EventBus:     eventBus,
	}, nil
}

func GetGameStateEventByCode(c context.Context, code string) (mahjongStruct.GameStateUpdated, error) {
	DB := database.DB
	queries := sqlc.New(DB)

	room, err := queries.GetRoomByCode(c, code)
	if err != nil {
		log.Info("Error getting room by code\n", err)
		return mahjongStruct.GameStateUpdated{}, fmt.Errorf("Unable to get room with code", err)
	}

	gameState, err := queries.GetGameStateById(c, room.GameStateID)
	if err != nil {
		log.Info("Error getting game state by id\n", err)
		return mahjongStruct.GameStateUpdated{}, fmt.Errorf("Unable to get room with code", err)
	}

	players, err := queries.GetPlayersByRoomId(c, room.ID)
	if err != nil {
		log.Info("Error getting players by room id\n", err)
		return mahjongStruct.GameStateUpdated{}, fmt.Errorf("Unable to get room with code", err)
	}

	return mahjongStruct.GameStateUpdated{
		RoomCode:  code,
		GameState: gameState,
		Players:   players,
	}, nil
}
