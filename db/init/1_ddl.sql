# CREATE TABLE IF NOT EXISTS `practice`. users (
CREATE SCHEMA IF NOT EXISTS `go_practice` DEFAULT CHARACTER SET utf8 ;
USE `go_practice`;

CREATE TABLE IF NOT EXISTS `go_practice`. users (
    number INT NOT NULL AUTO_INCREMENT UNIQUE,
    id VARCHAR(32) NOT NULL,
    name VARCHAR(32) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = Innodb
COMMENT = 'ユーザー';

CREATE TABLE IF NOT EXISTS `go_practice`. BeaconList (
    major VARCHAR(32) NOT NULL,
    beacon_number INT NOT NULL,
    PRIMARY KEY (major)
) ENGINE = Innodb
COMMENT = 'ビーコンの情報について';