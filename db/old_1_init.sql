CREATE DATABASE IF NOT EXISTS ca_mission;

USE ca_mission;

CREATE TABLE IF NOT EXISTS user (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  token VARCHAR(100) NOT NULL,
  created_at timestamp NOT NULL default current_timestamp,
  updated_at timestamp NOT NULL default current_timestamp on update current_timestamp
);

CREATE TABLE IF NOT EXISTS rarity (
  id int NOT NULL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  probability float NOT NULL
);

CREATE TABLE IF NOT EXISTS chara (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  rarity_id int NOT NULL,
  FOREIGN KEY (rarity_id) REFERENCES rarity(id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS user_chara_possession (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id int NOT NULL,
  chara_id int NOT NULL,
  FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT ON UPDATE CASCADE,
  FOREIGN KEY (chara_id) REFERENCES chara(id) ON DELETE RESTRICT ON UPDATE CASCADE
);

INSERT INTO
  rarity
VALUES
  (1, "SSR", 0.01),
  (2, "SR", 0.1),
  (3, "R", 0.89);

INSERT INTO
  chara (name, rarity_id)
VALUES
  ("めちゃめちゃレア1", 1),
  ("めちゃめちゃレア２", 1),
  ("めちゃレア１", 2),
  ("めちゃレア2", 2),
  ("めちゃレア3", 2),
  ("めちゃレア4", 2),
  ("レア1", 3),
  ("レア2", 3),
  ("レア3", 3),
  ("レア4", 3),
  ("レア5", 3),
  ("レア6", 3),
  ("レア7", 3),
  ("レア8", 3);