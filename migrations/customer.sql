-- customer.customer definition

CREATE TABLE `customer` (
                            `id` int(11) NOT NULL AUTO_INCREMENT,
                            `phone` varchar(100) NOT NULL,
                            `name` varchar(100) NOT NULL,
                            `balance` decimal(11,4) NOT NULL DEFAULT 0.0000,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;