CREATE DATABASE IF NOT EXISTS ca_mission;

USE ca_mission;

CREATE TABLE IF NOT EXISTS user (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  token VARCHAR(100) NOT NULL,
  created_at timestamp NOT NULL default current_timestamp,
  updated_at timestamp NOT NULL default current_timestamp on update current_timestamp
) DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS chara (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL
) DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS rarity (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  probability int NOT NULL
) DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS chara_rarity (
  chara_id int,
  rarity_id int,
  PRIMARY KEY(chara_id, rarity_id),
  FOREIGN KEY (chara_id) REFERENCES chara(id) ON DELETE CASCADE,
  FOREIGN KEY (rarity_id) REFERENCES rarity(id) ON DELETE CASCADE
) DEFAULT CHARSET = utf8;

INSERT INTO
  user (name, token)
VALUES
  ("kenshin", "help");