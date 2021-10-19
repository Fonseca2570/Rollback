CREATE DATABASE IF NOT EXISTS FirstName;
USE FirstName;

CREATE TABLE migrations
(
    name       varchar(50),
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE DATABASE IF NOT EXISTS LastName;
USE LastName;

CREATE TABLE migrations
(
    name       varchar(50),
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE DATABASE IF NOT EXISTS Email;
USE Email;

CREATE TABLE migrations
(
    name       varchar(50),
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);