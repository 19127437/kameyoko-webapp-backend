CREATE TABLE `presentations`
(
    `id`          BIGINT AUTO_INCREMENT PRIMARY KEY,
    `name`        VARCHAR(50) NOT NULL,
    `owner`       BIGINT,
    `created_at`  DATETIME,
    `modified_at` DATETIME
) DEFAULT CHARSET = utf8mb4;

CREATE TABLE `question_categories`
(
    `id`   BIGINT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL
) DEFAULT CHARSET = utf8mb4;

CREATE TABLE `question_types`
(
    `id`               BIGINT AUTO_INCREMENT PRIMARY KEY,
    `name`             VARCHAR(100) NOT NULL,
    `question_cate_id` BIGINT
) DEFAULT CHARSET = utf8mb4;

CREATE TABLE `slides`
(
    `id`         BIGINT AUTO_INCREMENT PRIMARY KEY,
    `pres_id`    BIGINT,
    `slide_type` BIGINT
) DEFAULT CHARSET = utf8mb4;

ALTER TABLE `presentations`
    ADD CONSTRAINT `pres_users_id_fk` FOREIGN KEY (`owner`) REFERENCES `users` (`id`)
        ON DELETE SET NULL;

ALTER TABLE `question_types`
    ADD CONSTRAINT `ques_types_categories_id_fk` FOREIGN KEY (`question_cate_id`) REFERENCES `question_categories` (`id`)
        ON DELETE SET NULL;

ALTER TABLE `slides`
    ADD (
        CONSTRAINT `slides_pres_id_fk` FOREIGN KEY (`pres_id`) REFERENCES `presentations` (`id`)
            ON DELETE SET NULL,
        CONSTRAINT `slides_ques_types_id_fk` FOREIGN KEY (`slide_type`) REFERENCES `question_types` (`id`)
            ON DELETE SET NULL
        );