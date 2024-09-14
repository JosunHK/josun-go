CREATE TABLE mahjong_action_log(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    room_id BIGINT NOT NULL,
    player_id BIGINT NOT NULL,
    score_delta INT NOT NULL
)
