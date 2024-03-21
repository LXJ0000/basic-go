SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS`webook` DEFAULT CHARACTER SET utf8mb4;
USE `webook`;

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
    (1, 1, "root", "test@qq.com", "181", "$2a$10$hvhqvCaBxA22R1s0D.gMMePhf6WQS6TQvlH9JKu36BLJMFcxVwcRy",
     "root", "", "Hello World", "www.jannan.top", unix_timestamp(), unix_timestamp());


-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
                         `id` bigint NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) NULL DEFAULT NULL,
                         `updated_at` datetime(3) NULL DEFAULT NULL,
                         `parent_id` bigint NULL DEFAULT NULL,
                         `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                         `path` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                         `component` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                         `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                         `order_num` tinyint NULL DEFAULT NULL,
                         `redirect` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                         `catalogue` tinyint(1) NULL DEFAULT NULL,
                         `hidden` tinyint(1) NULL DEFAULT NULL,
                         `keep_alive` tinyint(1) NULL DEFAULT NULL,
                         `external` tinyint(1) NULL DEFAULT NULL,
                         `external_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                         PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 49 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (2, '2022-10-31 09:41:03.000', '2023-12-27 23:26:43.807', 0, '文章管理', '/article', 'Layout', 'ic:twotone-article', 1, '/article/list', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (3, '2022-10-31 09:41:03.000', '2023-12-24 23:33:34.013', 0, '消息管理', '/message', 'Layout', 'ic:twotone-email', 2, '/message/comment	', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (4, '2022-10-31 09:41:03.000', '2023-12-24 23:32:35.177', 0, '用户管理', '/user', 'Layout', 'ph:user-list-bold', 4, '/user/list', 0, 0, 0, 0, NULL);
INSERT INTO `menu` VALUES (5, '2022-10-31 09:41:03.000', '2023-12-24 23:32:34.788', 0, '系统管理', '/setting', 'Layout', 'ion:md-settings', 5, '/setting/website', 0, 0, 0, 0, NULL);
INSERT INTO `menu` VALUES (6, '2022-10-31 09:41:03.000', '2023-12-24 23:22:29.519', 2, '发布文章', 'write', '/article/write', 'icon-park-outline:write', 1, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (8, '2022-10-31 09:41:03.000', '2023-12-21 20:58:29.873', 2, '文章列表', 'list', '/article/list', 'material-symbols:format-list-bulleted', 2, '', 0, 0, 0, 0, NULL);
INSERT INTO `menu` VALUES (9, '2022-10-31 09:41:03.000', '2022-11-01 01:18:30.931', 2, '分类管理', 'category', '/article/category', 'tabler:category', 3, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (10, '2022-10-31 09:41:03.000', '2022-11-01 01:18:35.502', 2, '标签管理', 'tag', '/article/tag', 'tabler:tag', 4, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (16, '2022-10-31 09:41:03.000', '2022-11-01 10:11:23.195', 0, '权限管理', '/auth', 'Layout', 'cib:adguard', 3, '/auth/menu', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (17, '2022-10-31 09:41:03.000', NULL, 16, '菜单管理', 'menu', '/auth/menu', 'ic:twotone-menu-book', 1, NULL, 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (23, '2022-10-31 09:41:03.000', NULL, 16, '接口管理', 'resource', '/auth/resource', 'mdi:api', 2, NULL, 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (24, '2022-10-31 09:41:03.000', '2022-10-31 10:09:18.913', 16, '角色管理', 'role', '/auth/role', 'carbon:user-role', 3, NULL, 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (25, '2022-10-31 10:11:09.232', '2022-11-01 01:29:48.520', 3, '评论管理', 'comment', '/message/comment', 'ic:twotone-comment', 1, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (26, '2022-10-31 10:12:01.546', '2022-11-01 01:29:54.130', 3, '留言管理', 'leave-msg', '/message/leave-msg', 'ic:twotone-message', 2, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (27, '2022-10-31 10:54:03.201', '2022-11-01 01:30:06.901', 4, '用户列表', 'list', '/user/list', 'mdi:account', 1, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (28, '2022-10-31 10:54:34.167', '2022-11-01 01:30:13.400', 4, '在线用户', 'online', '/user/online', 'ic:outline-online-prediction', 2, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (29, '2022-10-31 10:59:33.255', '2022-11-01 01:30:20.688', 5, '网站管理', 'website', '/setting/website', 'el:website', 1, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (30, '2022-10-31 11:00:09.997', '2022-11-01 01:30:24.097', 5, '页面管理', 'page', '/setting/page', 'iconoir:journal-page', 2, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (31, '2022-10-31 11:00:33.543', '2022-11-01 01:30:28.497', 5, '友链管理', 'link', '/setting/link', 'mdi:telegram', 3, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (32, '2022-10-31 11:01:00.444', '2022-11-01 01:30:33.186', 5, '关于我', 'about', '/setting/about', 'cib:about-me', 4, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (33, '2022-11-01 01:43:10.142', '2023-12-27 23:26:41.553', 0, '首页', '/home', '/home', 'ic:sharp-home', 0, '', 1, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (34, '2022-11-01 09:54:36.252', '2022-11-01 10:07:00.254', 2, '修改文章', 'write/:id', '/article/write', 'icon-park-outline:write', 1, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (36, '2022-11-04 15:50:45.993', '2023-12-24 23:32:33.538', 0, '日志管理', '/log', 'Layout', 'material-symbols:receipt-long-outline-rounded', 6, '/log/operation', 0, 0, 0, 0, NULL);
INSERT INTO `menu` VALUES (37, '2022-11-04 15:53:00.251', '2023-12-24 23:15:22.034', 36, '操作日志', 'operation', '/log/operation', 'mdi:book-open-page-variant-outline', 1, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (38, '2022-11-04 16:02:42.306', '2022-11-04 16:05:35.761', 36, '登录日志', 'login', '/log/login', 'material-symbols:login', 2, '', 0, 0, 1, 0, NULL);
INSERT INTO `menu` VALUES (39, '2022-12-07 20:47:08.349', '2023-12-24 23:33:35.701', 0, '个人中心', '/profile', '/profile', 'mdi:account', 7, '', 1, 0, 0, 0, NULL);
INSERT INTO `menu` VALUES (47, '2023-12-24 20:26:14.173', '2023-12-24 23:33:36.247', 0, '测试一级菜单', '/testone', 'Layout', '', 88, '', 0, 0, 0, 1, NULL);
INSERT INTO `menu` VALUES (48, '2023-12-24 23:26:19.441', '2023-12-24 23:26:27.704', 0, '测试外链', 'https://www.baidu.com', 'Layout', 'mdi-fan-speed-3', 66, '', 1, 0, 0, 1, '');