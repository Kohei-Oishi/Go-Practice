CREATE SCHEMA IF NOT EXISTS `go_practice` DEFAULT CHARACTER SET utf8 ;
USE `go_practice` ;

CREATE TABLE IF NOT EXISTS `go_practice`.`users` (
    `number` INT NOT NULL AUTO_INCREMENT,
    `id` VARCHAR(32) NOT NULL,
    `name` VARCHAR(32) NOT NULL,
    PRIMARY KEY (`id`)
);
-- ENGINE = InnoDB;
