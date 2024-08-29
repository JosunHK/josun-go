CREATE TABLE mahjong_player(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    room_id BIGINT NOT NULL,
    score INT NOT NULL,
    wind varchar(50) NOT NULL
)
