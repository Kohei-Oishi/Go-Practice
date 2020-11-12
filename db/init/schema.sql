CREATE SCHEMA IF NOT EXISTS `practice` DEFAULT CHARACTER SET utf8 ;
USE `practice` ;

CREATE TABLE IF NOT EXISTS `practice`. `users` (
    number INT NOT NULL AUTO_INCREMENT,
    id VARCHAR(32) NOT NULL,
    name VARCHAR(32) NOT NULL,
    PRIMARY KEY (id)
)ENGINE = Innodb,
DEFAULT character set utf8,
COMMENT = 'ユーザー';
