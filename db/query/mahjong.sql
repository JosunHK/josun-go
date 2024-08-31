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

-- name: GetRoomWithCode :one
SELECT * FROM mahjong_room 
WHERE room_code = ? LIMIT 1;
