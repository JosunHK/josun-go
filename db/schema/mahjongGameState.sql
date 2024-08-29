CREATE TABLE mahjong_game_state(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    room_id BIGINT NOT NULL,
    wind varchar(50) NOT NULL,
    round INT NOT NULL
)
