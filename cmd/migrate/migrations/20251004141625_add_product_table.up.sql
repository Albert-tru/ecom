CREATE TABLE IF NOT EXISTS products (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT,
    `image` VARCHAR(255),
    `price` DECIMAL(10, 2) NOT NULL,
    `quantity` INT NOT NULL DEFAULT 0,
    `createdat` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (`id`)
);