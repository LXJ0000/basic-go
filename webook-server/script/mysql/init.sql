create database webook;

use webook;

CREATE TABLE `user` (
                        `id` bigint NOT NULL AUTO_INCREMENT,
                        `user_id` bigint DEFAULT NULL,
                        `user_name` varchar(191) DEFAULT NULL,
                        `email` varchar(191) DEFAULT NULL,
                        `phone` varchar(191) DEFAULT NULL,
                        `password` longtext,
                        `nick_name` longtext,
                        `avatar` longtext,
                        `intro` longtext,
                        `web_site` longtext,
                        `create_at` bigint DEFAULT NULL,
                        `update_at` bigint DEFAULT NULL,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `uni_user_user_id` (`user_id`),
                        UNIQUE KEY `uni_user_user_name` (`user_name`),
                        UNIQUE KEY `uni_user_email` (`email`),
                        UNIQUE KEY `uni_user_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



insert into user(id, user_id, user_name, email, phone, password, nick_name, avatar, intro, web_site, create_at, update_at) values
    (1, 1, "root", "123@qq.com", "181", "$2a$10$hvhqvCaBxA22R1s0D.gMMePhf6WQS6TQvlH9JKu36BLJMFcxVwcRy",
     "root", "", "Hello World", "www.jannan.top", unix_timestamp(), unix_timestamp());