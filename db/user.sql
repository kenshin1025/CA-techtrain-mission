CREATE TABLE `user` (
          `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
          `name` VARCHAR(50) NOT NULL,
          `token` VARCHAR(200) NOT NULL,
          `created_at` timestamp NOT NULL default current_timestamp,
          `updated_at` timestamp NOT NULL default current_timestamp on update current_timestamp
          );