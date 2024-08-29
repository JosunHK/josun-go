CREATE TABLE mahjong_room(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    game_state_id BIGINT NOT NULL,
    owner_id BIGINT NOT NULL,
    room_code varchar(20) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    game_length varchar(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)
