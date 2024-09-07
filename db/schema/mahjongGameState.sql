CREATE TABLE mahjong_game_state(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    round_wind ENUM ('east', 'south', 'west', 'north') NOT NULL,
    seat_wind ENUM ('east', 'south', 'west', 'north') NOT NULL,
    round INT NOT NULL,
    kyoutaku INT NOT NULL DEFAULT 0,
    ended bool NOT NULL DEFAULT FALSE
)
