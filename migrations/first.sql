drop table if exists users;

CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email varchar(255),
    phrase varchar(255),
    flavor varchar(255),
    fans int NOT NULL default 0,
    videos int NOT NULL default 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY unique_email (email)
) ENGINE InnoDB;

