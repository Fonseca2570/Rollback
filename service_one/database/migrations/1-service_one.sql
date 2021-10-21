CREATE TABLE user_one
(
    id         int         NOT NULL PRIMARY KEY,
    id_users   int         NOT NULL,
    first_name varchar(50) NOT NULL
);

INSERT INTO user_one (id, id_users, first_name)
VALUES (1, 100, 'google');

CREATE TABLE `user_one_history`
(
    `id`           INT NOT NULL AUTO_INCREMENT,
    `user_one_id`  INT,
    `id_users`     int NOT NULL,
    `first_name`   VARCHAR(50) NULL,
    `date_changed` DATETIME,
    PRIMARY KEY (`id`)
);