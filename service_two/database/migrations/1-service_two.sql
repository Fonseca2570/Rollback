CREATE TABLE user_two
(
    id         int         NOT NULL PRIMARY KEY,
    id_users   int         NOT NULL,
    last_name varchar(50) NOT NULL
);

INSERT INTO user_two (id, id_users, last_name)
VALUES (1, 100, 'www');

CREATE TABLE `user_two_history`
(
    `id`           INT NOT NULL AUTO_INCREMENT,
    `user_two_id`  INT,
    `id_users`     int NOT NULL,
    `last_name`   VARCHAR(50) NULL,
    `date_changed` DATETIME,
    PRIMARY KEY (`id`)
);