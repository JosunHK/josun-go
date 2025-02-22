-- name: CreateMahjongRoom :execresult
INSERT INTO mahjong_room(
    game_state_id,
    owner_Id,
    room_code, 
    game_length
) VALUES (
    ?, ?, ?, ?
);

-- name: CreateActionLog :execresult
INSERT INTO mahjong_action_log(
    room_id,
    player_id,
    score_delta
) VALUES (
    ?, ?, ?
);

-- name: CreateMahjongRoomOwner :execresult
INSERT INTO mahjong_room_owner(
    user_id,
    guest_id
) VALUES (
    ?, ?
);

-- name: CreateMahjongGameState :execresult
INSERT INTO mahjong_game_state(
    round_wind,
    seat_wind,
    round 
) VALUES (
    ?, ?, ? 
);

-- name: CreateMahjongPlayer :execresult
INSERT INTO mahjong_player(
    room_id,
    name,
    score,
    wind
) VALUES (
    ?, ?, ?, ?
);

-- name: GetRoomByCode :one
SELECT * FROM mahjong_room 
WHERE room_code = ? LIMIT 1;

-- name: GetRoomById :one
SELECT * FROM mahjong_room 
WHERE id = ? LIMIT 1;

-- name: GetGameStateById :one
SELECT * FROM mahjong_game_state 
WHERE id = ? LIMIT 1;

-- name: GetPlayersByRoomId :many
SELECT * FROM mahjong_player
WHERE room_id = ?;

-- name: GetOwnerById :one
SELECT * FROM mahjong_room_owner
WHERE id = ? LIMIT 1;

-- name: GetOwnerByUUIDorUserId :one
SELECT * FROM mahjong_room_owner
WHERE (user_id = ? 
OR guest_id = ?) LIMIT 1;

-- name: GetPlayerById :one
SELECT * FROM mahjong_player
WHERE id = ? LIMIT 1;

-- name: GetGameStateByRoomCode :one
SELECT * FROM mahjong_game_state
WHERE EXISTS (
    SELECT * FROM mahjong_room 
    WHERE room_code = ? 
    AND active = TRUE
    AND mahjong_game_state.id = mahjong_room.game_state_id
) LIMIT 1;

-- name: GetPlayersByRoomCode :many
SELECT * FROM mahjong_player
WHERE EXISTS (
    SELECT * FROM mahjong_room
    WHERE room_code = ?
    AND active = TRUE
    AND mahjong_player.room_id = mahjong_room.id
);

-- name: GetPlayerCountByRoomId :one
SELECT COUNT(*) FROM mahjong_player
WHERE id IN (sqlc.slice(ids)) AND room_id = ?;

-- name: UpdatePlayerScore :exec
UPDATE mahjong_player 
SET score = ? 
WHERE id = ?;

-- name: UpdateGameState :exec
UPDATE mahjong_game_state 
SET round_wind = ?,
    seat_wind = ?,
    round = ?,
    kyoutaku = ?,
    ended = ?
WHERE id = ?;

-- name: EndGameByCode :exec
UPDATE mahjong_room
SET active = false
WHERE room_code = ?;
