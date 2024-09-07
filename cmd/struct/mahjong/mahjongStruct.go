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

type DrawPlayer struct {
	PlayerId int64 `schema:"playerId,required"`
	Tenpai   bool  `schema:"tenpai,required"`
}

type GameData struct {
	Room      sqlc.MahjongRoom
	GameState sqlc.MahjongGameState
	Players   []sqlc.MahjongPlayer
}
