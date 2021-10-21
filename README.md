# Roolback
Trying too rollback between various services dinamic


protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service1.proto

https://stackoverflow.com/questions/6787794/how-to-log-all-changes-in-a-mysql-table-to-a-second-one/23407137


CREATE TABLE `house` (
`id` INT NOT NULL AUTO_INCREMENT,
`name` VARCHAR(50) NULL,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB;

CREATE TABLE `house_history` (
`id` INT NOT NULL AUTO_INCREMENT,
`house_id` INT ,
`name` VARCHAR(50) NULL,
`date_changed` DATETIME,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB;

DELIMITER //
CREATE TRIGGER `updateHouse` BEFORE UPDATE ON `house`
FOR EACH ROW INSERT INTO house_history
SELECT NULL, h.*, NOW() FROM house h WHERE id = OLD.id
//
DELIMITER ;