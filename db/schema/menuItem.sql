CREATE TABLE menu_item (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    menu_collection_id BIGINT NOT NULL,
    label varchar(150) NOT NULL,
    value varchar(150) NOT NULL,
    sort_order INT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE
);
