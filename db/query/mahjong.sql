-- name: CreateMahjongRoom :execresult
INSERT INTO mahjong_room(
    game_state_id,
    owner_Id,
    room_code, 
    game_length
) VALUES (
    ?, ?, ?, ?
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

-- name: GetGameStateById :one
SELECT * FROM mahjong_game_state 
WHERE id = ? LIMIT 1;

-- name: GetPlayersByRoomId :many
SELECT * FROM mahjong_player
WHERE room_id = ? LIMIT 1;

-- name: GetOwnerById :one
SELECT * FROM mahjong_room_owner
WHERE id = ? LIMIT 1;

-- name: GetOwnerByUUIDorUserId :one
SELECT * FROM mahjong_room_owner
WHERE  user_id = ? 
OR guest_id = ? LIMIT 1;

-- name: GetPlayerById :one
SELECT * FROM mahjong_player
WHERE id = ? LIMIT 1;

-- name: GetPlayersByRoomCode :many
SELECT * FROM mahjong_player
WHERE EXISTS (
    SELECT * FROM mahjong_room
    WHERE room_code = ?
    AND mahjong_player.room_id = mahjong_room.id
);

-- name: GetPlayerCountByRoomId :one
SELECT COUNT(*) FROM MAHJONG_PLAYER 
WHERE id IN (sqlc.slice(ids)) AND room_id = ?;

-- name: UpdatePlayerScore :exec
UPDATE mahjong_player 
SET score = ? 
WHERE id = ?;
