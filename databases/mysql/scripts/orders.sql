START TRANSACTION;

USE exampledb;

CREATE TABLE orders (
    id int UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    orderNumber VARCHAR(255) NOT NULL UNIQUE,
    currencyCode varchar(255),
    amount float,
    created_at timestamp,
    updated_at timestamp
);

INSERT INTO orders (orderNumber, currencyCode, amount, created_at, updated_at) 
VALUES 
('ORD001', 'USD', 100.50, NOW(), NOW()),
('ORD002', 'EUR', 75.25, NOW(), NOW()),
('ORD003', 'GBP', 150.80, NOW(), NOW()),
('ORD004', 'JPY', 2000.00, NOW(), NOW()),
('ORD005', 'CAD', 90.75, NOW(), NOW());

COMMIT;