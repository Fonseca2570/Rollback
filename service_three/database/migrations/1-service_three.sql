CREATE TABLE user_three
(
    id       int         NOT NULL PRIMARY KEY,
    id_users int         NOT NULL,
    email    varchar(50) NOT NULL
);

INSERT INTO user_three (id, id_users, email)
VALUES (1, 100, 'www');

CREATE TABLE `user_three_history`
(
    `id`            INT NOT NULL AUTO_INCREMENT,
    `user_three_id` INT,
    `id_users`      int NOT NULL,
    `email`         VARCHAR(50) NULL,
    `date_changed`  DATETIME,
    PRIMARY KEY (`id`)
);