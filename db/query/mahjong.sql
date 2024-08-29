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
