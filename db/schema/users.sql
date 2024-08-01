CREATE TABLE users (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(150) NOT NULL,
  email varchar(250) NOT NULL UNIQUE,
  password varchar(500) NOT NULL
);
