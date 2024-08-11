CREATE TABLE menuItem (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    menuCollectionId BIGINT NOT NULL,
    label varchar(150) NOT NULL,
    value varchar(150) NOT NULL,
    sortOrder INT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE
);
