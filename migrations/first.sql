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

drop table if exists videos;

CREATE TABLE videos (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title varchar(255),
    user_id int,
    comments int not null default 0,
    status varchar(255),
    url_safe_name varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    KEY vbyu (user_id),
    UNIQUE key urlsafe (url_safe_name)
) ENGINE InnoDB;

