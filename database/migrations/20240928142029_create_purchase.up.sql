CREATE TABLE `purchase_headers` (
    `id` VARCHAR(100) NOT NULL,
    `purchase_date` DATE NOT NULL,
    `status` varchar(10) NOT NULL,
    `supplier_id` bigint unsigned NOT NULL,
    `branch_id` bigint unsigned NOT NULL,
    `payment_method_id` bigint unsigned NOT NULL,
    `shipping_method_id` bigint unsigned NOT NULL,
    `shipping_amount` DECIMAL(10,2),
    `discount_amount` DECIMAL(10,2),
    `total_amount` DECIMAL(10,2),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`supplier_id`) REFERENCES `suppliers`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`branch_id`) REFERENCES `branches`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`payment_method_id`) REFERENCES `payment_method`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`shipping_method_id`) REFERENCES `shipping_method`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE `purchase_lines` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `purchase_id` VARCHAR(100),
    `item_id` bigint unsigned,
    `item_price` DECIMAL(10,2),
    `item_discount` DECIMAL(10,2),
    `item_qty` DECIMAL(10,2),
    `item_total` DECIMAL(10,2),    
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`purchase_id`) REFERENCES `purchase_headers`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`item_id`) REFERENCES `products`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
