package mahjongStruct

import (
	sqlc "github.com/JosunHK/josun-go.git/db/generated"
)

type GameUpdated struct {
	RoomCode string `json:"room_id"`
}

type GameStateUpdated struct {
	RoomCode  string                `json:"room_code"`
	RoomID    int64                 `json:"room_id"`
	GameState sqlc.MahjongGameState `json:"game_state"`
	Players   []sqlc.MahjongPlayer  `json:"players"`
}

type DrawPlayer struct {
	PlayerId int64 `schema:"playerId,required"`
	NoTen    bool  `schema:"noTen,required"`
}

type RiichiPlayer struct {
	PlayerId int64 `schema:"playerId,required"`
	Riichi   bool  `schema:"riichi,required"`
}

type GameData struct {
	Room      sqlc.MahjongRoom
	GameState sqlc.MahjongGameState
	Players   []sqlc.MahjongPlayer
}

type Winner struct {
	PlayerId int64 `schema:"playerId"`
	Han      int   `schema:"han"`
	Fu       int   `schema:"fu"`
}

type WinForm struct {
	UpdateType    string         `schema:"updateType,required"`
	Winners       []Winner       `schema:"winner"`
	IsTsumo       bool           `schema:"isTsumo,required"`
	LoserId       int64          `schema:"loserId"`
	RiichiPlayers []RiichiPlayer `schema:"riichiPlayers,required"`
}

type DrawForm struct {
	UpdateType    string         `schema:"updateType,required"`
	DrawPlayers   []DrawPlayer   `schema:"drawPlayers,required"`
	RiichiPlayers []RiichiPlayer `schema:"riichiPlayers,required"`
}

type ManualForm struct {
	UpdateType string `schema:"updateType,required"`
	PlayerId   int64  `schema:"playerId,required"`
	Score      int    `schema:"score,default:0"`
}

type Score struct {
	OyaTsumo   int `json:"oyaTsumo"`
	KoTsumoKo  int `json:"koTsumoKo"`
	KoTsumoOya int `json:"koTsumoOya"`
	OyaDirect  int `json:"oyaDirect"`
	KoDirect   int `json:"koDirect"`
}

var ScoreMap = map[int]map[int]Score{
	1: {
		20:  Score{OyaTsumo: 0, KoTsumoKo: 0, KoTsumoOya: 0, OyaDirect: 0, KoDirect: 0},
		25:  Score{OyaTsumo: 0, KoTsumoKo: 0, KoTsumoOya: 0, OyaDirect: 0, KoDirect: 0},
		30:  Score{OyaTsumo: 500, KoTsumoKo: 300, KoTsumoOya: 500, OyaDirect: 1500, KoDirect: 1000},
		40:  Score{OyaTsumo: 700, KoTsumoKo: 400, KoTsumoOya: 700, OyaDirect: 2000, KoDirect: 1300},
		50:  Score{OyaTsumo: 800, KoTsumoKo: 400, KoTsumoOya: 800, OyaDirect: 2400, KoDirect: 1600},
		60:  Score{OyaTsumo: 1000, KoTsumoKo: 500, KoTsumoOya: 1000, OyaDirect: 2900, KoDirect: 2000},
		70:  Score{OyaTsumo: 1200, KoTsumoKo: 600, KoTsumoOya: 1200, OyaDirect: 3400, KoDirect: 2300},
		80:  Score{OyaTsumo: 1300, KoTsumoKo: 700, KoTsumoOya: 1300, OyaDirect: 3900, KoDirect: 2600},
		90:  Score{OyaTsumo: 1500, KoTsumoKo: 800, KoTsumoOya: 1500, OyaDirect: 4400, KoDirect: 2900},
		100: Score{OyaTsumo: 1600, KoTsumoKo: 800, KoTsumoOya: 1600, OyaDirect: 4800, KoDirect: 3200},
		110: Score{OyaTsumo: 1800, KoTsumoKo: 900, KoTsumoOya: 1800, OyaDirect: 5300, KoDirect: 3600},
	},
	2: {
		20:  Score{OyaTsumo: 700, KoTsumoKo: 400, KoTsumoOya: 700, OyaDirect: 0, KoDirect: 0},
		25:  Score{OyaTsumo: 800, KoTsumoKo: 400, KoTsumoOya: 800, OyaDirect: 2400, KoDirect: 1600},
		30:  Score{OyaTsumo: 1000, KoTsumoKo: 500, KoTsumoOya: 1000, OyaDirect: 2900, KoDirect: 2000},
		40:  Score{OyaTsumo: 1300, KoTsumoKo: 700, KoTsumoOya: 1300, OyaDirect: 3900, KoDirect: 2600},
		50:  Score{OyaTsumo: 1600, KoTsumoKo: 800, KoTsumoOya: 1600, OyaDirect: 4800, KoDirect: 3200},
		60:  Score{OyaTsumo: 1000, KoTsumoKo: 1000, KoTsumoOya: 2000, OyaDirect: 5800, KoDirect: 3900},
		70:  Score{OyaTsumo: 1200, KoTsumoKo: 1200, KoTsumoOya: 2300, OyaDirect: 6800, KoDirect: 4500},
		80:  Score{OyaTsumo: 1300, KoTsumoKo: 1300, KoTsumoOya: 2600, OyaDirect: 7700, KoDirect: 5200},
		90:  Score{OyaTsumo: 1500, KoTsumoKo: 1500, KoTsumoOya: 2900, OyaDirect: 8700, KoDirect: 5800},
		100: Score{OyaTsumo: 1600, KoTsumoKo: 1600, KoTsumoOya: 3200, OyaDirect: 9600, KoDirect: 6400},
		110: Score{OyaTsumo: 3600, KoTsumoKo: 1800, KoTsumoOya: 3600, OyaDirect: 10600, KoDirect: 7100},
	},
	3: {
		20:  Score{OyaTsumo: 1300, KoTsumoKo: 700, KoTsumoOya: 1300, OyaDirect: 0, KoDirect: 0},
		25:  Score{OyaTsumo: 1600, KoTsumoKo: 800, KoTsumoOya: 1600, OyaDirect: 4800, KoDirect: 3200},
		30:  Score{OyaTsumo: 2000, KoTsumoKo: 1000, KoTsumoOya: 2000, OyaDirect: 5800, KoDirect: 3900},
		40:  Score{OyaTsumo: 2600, KoTsumoKo: 1300, KoTsumoOya: 2600, OyaDirect: 7700, KoDirect: 5200},
		50:  Score{OyaTsumo: 3200, KoTsumoKo: 1600, KoTsumoOya: 3200, OyaDirect: 9600, KoDirect: 6400},
		60:  Score{OyaTsumo: 3900, KoTsumoKo: 2000, KoTsumoOya: 3900, OyaDirect: 11600, KoDirect: 7700},
		70:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		80:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		90:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		100: Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		110: Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
	},
	4: {
		20:  Score{OyaTsumo: 2600, KoTsumoKo: 1300, KoTsumoOya: 2600, OyaDirect: 0, KoDirect: 0},
		25:  Score{OyaTsumo: 3200, KoTsumoKo: 1600, KoTsumoOya: 3200, OyaDirect: 9600, KoDirect: 6400},
		30:  Score{OyaTsumo: 3900, KoTsumoKo: 2000, KoTsumoOya: 3900, OyaDirect: 11600, KoDirect: 7700},
		40:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		50:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		60:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		70:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		80:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		90:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		100: Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		110: Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
	},
	5: {
		20:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		25:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		30:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		40:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		50:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		60:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		70:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		80:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		90:  Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		100: Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
		110: Score{OyaTsumo: 4000, KoTsumoKo: 2000, KoTsumoOya: 4000, OyaDirect: 12000, KoDirect: 8000},
	},
	6: {
		20:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		25:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		30:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		40:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		50:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		60:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		70:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		80:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		90:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		100: Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		110: Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
	},
	7: {
		20:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		25:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		30:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		40:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		50:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		60:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		70:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		80:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		90:  Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		100: Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
		110: Score{OyaTsumo: 6000, KoTsumoKo: 3000, KoTsumoOya: 6000, OyaDirect: 18000, KoDirect: 12000},
	},
	8: {
		20:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		25:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		30:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		40:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		50:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		60:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		70:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		80:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		90:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		100: Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		110: Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
	},
	9: {
		20:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		25:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		30:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		40:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		50:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		60:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		70:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		80:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		90:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		100: Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		110: Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
	},
	10: {
		20:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		25:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		30:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		40:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		50:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		60:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		70:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		80:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		90:  Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		100: Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
		110: Score{OyaTsumo: 8000, KoTsumoKo: 4000, KoTsumoOya: 8000, OyaDirect: 24000, KoDirect: 16000},
	},
	11: {
		20:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		25:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		30:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		40:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		50:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		60:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		70:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		80:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		90:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		100: Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		110: Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
	},
	12: {
		20:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		25:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		30:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		40:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		50:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		60:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		70:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		80:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		90:  Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		100: Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
		110: Score{OyaTsumo: 12000, KoTsumoKo: 6000, KoTsumoOya: 12000, OyaDirect: 36000, KoDirect: 24000},
	},
	13: {
		20:  Score{OyaTsumo: 16000, KoTsumoKo: 8000, KoTsumoOya: 16000, OyaDirect: 48000, KoDirect: 32000},
		25:  Score{OyaTsumo: 16000, KoTsumoKo: 8000, KoTsumoOya: 16000, OyaDirect: 48000, KoDirect: 32000},
		30:  Score{OyaTsumo: 16000, KoTsumoKo: 8000, KoTsumoOya: 16000, OyaDirect: 48000, KoDirect: 32000},
		40:  Score{OyaTsumo: 16000, KoTsumoKo: 8000, KoTsumoOya: 16000, OyaDirect: 48000, KoDirect: 32000},
		50:  Score{OyaTsumo: 16000, KoTsumoKo: 8000, KoTsumoOya: 16000, OyaDirect: 48000, KoDirect: 32000},
		60:  Score{OyaTsumo: 16000, KoTsumoKo: 8000, KoTsumoOya: 16000, OyaDirect: 48000, KoDirect: 32000},
		70:  Score{OyaTsumo: 16000, KoTsumoKo: 8000, KoTsumoOya: 16000, OyaDirect: 48000, KoDirect: 32000},
		80:  Score{OyaTsumo: 16000, KoTsumoKo: 8000, KoTsumoOya: 16000, OyaDirect: 48000, KoDirect: 32000},
		90:  Score{OyaTsumo: 16000, KoTsumoKo: 8000, KoTsumoOya: 16000, OyaDirect: 48000, KoDirect: 32000},
		100: Score{OyaTsumo: 16000, KoTsumoKo: 8000, KoTsumoOya: 16000, OyaDirect: 48000, KoDirect: 32000},
		110: Score{OyaTsumo: 16000, KoTsumoKo: 8000, KoTsumoOya: 16000, OyaDirect: 48000, KoDirect: 32000},
	},
	26: {
		20:  Score{OyaTsumo: 32000, KoTsumoKo: 16000, KoTsumoOya: 32000, OyaDirect: 96000, KoDirect: 64000},
		25:  Score{OyaTsumo: 32000, KoTsumoKo: 16000, KoTsumoOya: 32000, OyaDirect: 96000, KoDirect: 64000},
		30:  Score{OyaTsumo: 32000, KoTsumoKo: 16000, KoTsumoOya: 32000, OyaDirect: 96000, KoDirect: 64000},
		40:  Score{OyaTsumo: 32000, KoTsumoKo: 16000, KoTsumoOya: 32000, OyaDirect: 96000, KoDirect: 64000},
		50:  Score{OyaTsumo: 32000, KoTsumoKo: 16000, KoTsumoOya: 32000, OyaDirect: 96000, KoDirect: 64000},
		60:  Score{OyaTsumo: 32000, KoTsumoKo: 16000, KoTsumoOya: 32000, OyaDirect: 96000, KoDirect: 64000},
		70:  Score{OyaTsumo: 32000, KoTsumoKo: 16000, KoTsumoOya: 32000, OyaDirect: 96000, KoDirect: 64000},
		80:  Score{OyaTsumo: 32000, KoTsumoKo: 16000, KoTsumoOya: 32000, OyaDirect: 96000, KoDirect: 64000},
		90:  Score{OyaTsumo: 32000, KoTsumoKo: 16000, KoTsumoOya: 32000, OyaDirect: 96000, KoDirect: 64000},
		100: Score{OyaTsumo: 32000, KoTsumoKo: 16000, KoTsumoOya: 32000, OyaDirect: 96000, KoDirect: 64000},
		110: Score{OyaTsumo: 32000, KoTsumoKo: 16000, KoTsumoOya: 32000, OyaDirect: 96000, KoDirect: 64000},
	},
	39: {
		20:  Score{OyaTsumo: 48000, KoTsumoKo: 24000, KoTsumoOya: 48000, OyaDirect: 144000, KoDirect: 96000},
		25:  Score{OyaTsumo: 48000, KoTsumoKo: 24000, KoTsumoOya: 48000, OyaDirect: 144000, KoDirect: 96000},
		30:  Score{OyaTsumo: 48000, KoTsumoKo: 24000, KoTsumoOya: 48000, OyaDirect: 144000, KoDirect: 96000},
		40:  Score{OyaTsumo: 48000, KoTsumoKo: 24000, KoTsumoOya: 48000, OyaDirect: 144000, KoDirect: 96000},
		50:  Score{OyaTsumo: 48000, KoTsumoKo: 24000, KoTsumoOya: 48000, OyaDirect: 144000, KoDirect: 96000},
		60:  Score{OyaTsumo: 48000, KoTsumoKo: 24000, KoTsumoOya: 48000, OyaDirect: 144000, KoDirect: 96000},
		70:  Score{OyaTsumo: 48000, KoTsumoKo: 24000, KoTsumoOya: 48000, OyaDirect: 144000, KoDirect: 96000},
		80:  Score{OyaTsumo: 48000, KoTsumoKo: 24000, KoTsumoOya: 48000, OyaDirect: 144000, KoDirect: 96000},
		90:  Score{OyaTsumo: 48000, KoTsumoKo: 24000, KoTsumoOya: 48000, OyaDirect: 144000, KoDirect: 96000},
		100: Score{OyaTsumo: 48000, KoTsumoKo: 24000, KoTsumoOya: 48000, OyaDirect: 144000, KoDirect: 96000},
		110: Score{OyaTsumo: 48000, KoTsumoKo: 24000, KoTsumoOya: 48000, OyaDirect: 144000, KoDirect: 96000},
	},
	52: {
		20:  Score{OyaTsumo: 64000, KoTsumoKo: 32000, KoTsumoOya: 64000, OyaDirect: 192000, KoDirect: 128000},
		25:  Score{OyaTsumo: 64000, KoTsumoKo: 32000, KoTsumoOya: 64000, OyaDirect: 192000, KoDirect: 128000},
		30:  Score{OyaTsumo: 64000, KoTsumoKo: 32000, KoTsumoOya: 64000, OyaDirect: 192000, KoDirect: 128000},
		40:  Score{OyaTsumo: 64000, KoTsumoKo: 32000, KoTsumoOya: 64000, OyaDirect: 192000, KoDirect: 128000},
		50:  Score{OyaTsumo: 64000, KoTsumoKo: 32000, KoTsumoOya: 64000, OyaDirect: 192000, KoDirect: 128000},
		60:  Score{OyaTsumo: 64000, KoTsumoKo: 32000, KoTsumoOya: 64000, OyaDirect: 192000, KoDirect: 128000},
		70:  Score{OyaTsumo: 64000, KoTsumoKo: 32000, KoTsumoOya: 64000, OyaDirect: 192000, KoDirect: 128000},
		80:  Score{OyaTsumo: 64000, KoTsumoKo: 32000, KoTsumoOya: 64000, OyaDirect: 192000, KoDirect: 128000},
		90:  Score{OyaTsumo: 64000, KoTsumoKo: 32000, KoTsumoOya: 64000, OyaDirect: 192000, KoDirect: 128000},
		100: Score{OyaTsumo: 64000, KoTsumoKo: 32000, KoTsumoOya: 64000, OyaDirect: 192000, KoDirect: 128000},
		110: Score{OyaTsumo: 64000, KoTsumoKo: 32000, KoTsumoOya: 64000, OyaDirect: 192000, KoDirect: 128000},
	},
	65: {
		20:  Score{OyaTsumo: 96000, KoTsumoKo: 48000, KoTsumoOya: 96000, OyaDirect: 288000, KoDirect: 192000},
		25:  Score{OyaTsumo: 96000, KoTsumoKo: 48000, KoTsumoOya: 96000, OyaDirect: 288000, KoDirect: 192000},
		30:  Score{OyaTsumo: 96000, KoTsumoKo: 48000, KoTsumoOya: 96000, OyaDirect: 288000, KoDirect: 192000},
		40:  Score{OyaTsumo: 96000, KoTsumoKo: 48000, KoTsumoOya: 96000, OyaDirect: 288000, KoDirect: 192000},
		50:  Score{OyaTsumo: 96000, KoTsumoKo: 48000, KoTsumoOya: 96000, OyaDirect: 288000, KoDirect: 192000},
		60:  Score{OyaTsumo: 96000, KoTsumoKo: 48000, KoTsumoOya: 96000, OyaDirect: 288000, KoDirect: 192000},
		70:  Score{OyaTsumo: 96000, KoTsumoKo: 48000, KoTsumoOya: 96000, OyaDirect: 288000, KoDirect: 192000},
		80:  Score{OyaTsumo: 96000, KoTsumoKo: 48000, KoTsumoOya: 96000, OyaDirect: 288000, KoDirect: 192000},
		90:  Score{OyaTsumo: 96000, KoTsumoKo: 48000, KoTsumoOya: 96000, OyaDirect: 288000, KoDirect: 192000},
		100: Score{OyaTsumo: 96000, KoTsumoKo: 48000, KoTsumoOya: 96000, OyaDirect: 288000, KoDirect: 192000},
		110: Score{OyaTsumo: 96000, KoTsumoKo: 48000, KoTsumoOya: 96000, OyaDirect: 288000, KoDirect: 192000},
	},
	78: {
		20:  Score{OyaTsumo: 128000, KoTsumoKo: 64000, KoTsumoOya: 128000, OyaDirect: 384000, KoDirect: 256000},
		25:  Score{OyaTsumo: 128000, KoTsumoKo: 64000, KoTsumoOya: 128000, OyaDirect: 384000, KoDirect: 256000},
		30:  Score{OyaTsumo: 128000, KoTsumoKo: 64000, KoTsumoOya: 128000, OyaDirect: 384000, KoDirect: 256000},
		40:  Score{OyaTsumo: 128000, KoTsumoKo: 64000, KoTsumoOya: 128000, OyaDirect: 384000, KoDirect: 256000},
		50:  Score{OyaTsumo: 128000, KoTsumoKo: 64000, KoTsumoOya: 128000, OyaDirect: 384000, KoDirect: 256000},
		60:  Score{OyaTsumo: 128000, KoTsumoKo: 64000, KoTsumoOya: 128000, OyaDirect: 384000, KoDirect: 256000},
		70:  Score{OyaTsumo: 128000, KoTsumoKo: 64000, KoTsumoOya: 128000, OyaDirect: 384000, KoDirect: 256000},
		80:  Score{OyaTsumo: 128000, KoTsumoKo: 64000, KoTsumoOya: 128000, OyaDirect: 384000, KoDirect: 256000},
		90:  Score{OyaTsumo: 128000, KoTsumoKo: 64000, KoTsumoOya: 128000, OyaDirect: 384000, KoDirect: 256000},
		100: Score{OyaTsumo: 128000, KoTsumoKo: 64000, KoTsumoOya: 128000, OyaDirect: 384000, KoDirect: 256000},
		110: Score{OyaTsumo: 128000, KoTsumoKo: 64000, KoTsumoOya: 128000, OyaDirect: 384000, KoDirect: 256000},
	},
}
