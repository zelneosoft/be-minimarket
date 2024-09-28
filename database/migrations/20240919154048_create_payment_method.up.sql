CREATE TABLE `payment_method` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) DEFAULT NULL,
    `is_active_for_sales` int,
    `is_active_for_purchase` int,
    `is_active` int,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

INSERT INTO `payment_method` (`name`, `is_active_for_sales`, `is_active_for_purchase`, `is_active`, `created_at`, `updated_at`) 
    VALUES  
    ('Credit Card', 1, 0, 1, NOW(), NOW()), 
    ('Bank Transfer', 1, 1, 1, NOW(), NOW()), 
    ('Cash', 1, 1, 1, NOW(), NOW()),
    ('E-Wallet', 1, 0, 1, NOW(), NOW());
