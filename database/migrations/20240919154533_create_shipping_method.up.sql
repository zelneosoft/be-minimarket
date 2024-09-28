CREATE TABLE `shipping_method` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) DEFAULT NULL,
    `is_active` int,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

INSERT INTO `shipping_method` (`name`, `is_active`, `created_at`, `updated_at`) 
    VALUES  
    ('Self Pickup', 1, NOW(), NOW()), 
    ('Delivery By Supplier', 1, NOW(), NOW());
