CREATE TABLE mahjong_player(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    roomId BIGINT NOT NULL,
    score INT NOT NULL,
    wind varchar(50) NOT NULL,
);
