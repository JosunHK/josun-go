package mahjongStruct

import (
	sqlc "github.com/JosunHK/josun-go.git/db/generated"
)

type GameUpdated struct {
	RoomCode string `json:"room_id"`
}

type GameStateUpdated struct {
	RoomCode  string                `json:"room_id"`
	GameState sqlc.MahjongGameState `json:"game_state"`
	Players   []sqlc.MahjongPlayer  `json:"players"`
}
