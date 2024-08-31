CREATE TABLE mahjong_player(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    room_id BIGINT NOT NULL,
    name varchar(50) NOT NULL,
    score INT NOT NULL,
    wind ENUM ('east', 'south', 'west', 'north') NOT NULL,
    unique key (room_id, wind),

    FOREIGN KEY (room_id) 
    REFERENCES mahjong_room(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
)
