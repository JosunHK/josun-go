CREATE TABLE mahjong_room(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    game_state_id BIGINT NOT NULL,
    owner_id BIGINT NOT NULL,
    room_code varchar(20) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    game_length ENUM ('han_chan', 'tonpuu') NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY (room_code ,(nullif(active, FALSE))),

    FOREIGN KEY (game_state_id) 
    REFERENCES mahjong_game_state(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,

    FOREIGN KEY (owner_id) 
    REFERENCES mahjong_room_owner(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
)
