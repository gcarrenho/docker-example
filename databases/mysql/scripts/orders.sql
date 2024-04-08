START TRANSACTION;

USE exampledb;

/*create table orders (
    id int UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(55) NOT NULL,
    age int
);*/

CREATE TABLE orders (
    id int UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    orderNumer VARCHAR(255) NOT NULL,
    currencyCode varchar(255),
    amount float,
    created_at timestamp,
    updated_at timestamp
);

COMMIT;