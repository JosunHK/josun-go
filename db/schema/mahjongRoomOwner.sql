CREATE TABLE mahjong_room_owner(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    guest_id char(36) NOT NULL
)
