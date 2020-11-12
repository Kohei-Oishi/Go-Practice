# CREATE TABLE IF NOT EXISTS `practice`. users (
CREATE SCHEMA IF NOT EXISTS `Practice` DEFAULT CHARACTER SET utf8 ;
USE `Practice`;

CREATE TABLE IF NOT EXISTS `Practice`. users (
    number INT NOT NULL AUTO_INCREMENT UNIQUE,
    id VARCHAR(32) NOT NULL,
    name VARCHAR(32) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = Innodb
COMMENT = 'ユーザー';

CREATE TABLE IF NOT EXISTS `Practice`. BeaconList (
    uuid VARCHAR(32) NOT NULL,
    major VARCHAR(32) NOT NULL,
    minor VARCHAR(32) NOT NULL,
    PRIMARY KEY (uuid)
) ENGINE = Innodb
    COMMENT = 'ユーザー';
