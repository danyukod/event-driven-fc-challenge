USE wallet;

CREATE TABLE client (
                        id VARCHAR(36) PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        email VARCHAR(255) NOT NULL UNIQUE,
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE account (
                         id VARCHAR(36) PRIMARY KEY,
                         client_id VARCHAR(36),
                         balance DECIMAL(10, 2) NOT NULL,
                         created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         FOREIGN KEY (client_id) REFERENCES client(id)
);

CREATE TABLE transaction (
                             id VARCHAR(36) PRIMARY KEY,
                             account_id_from VARCHAR(36),
                             account_id_to VARCHAR(36),
                             amount DECIMAL(10, 2) NOT NULL,
                             created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             FOREIGN KEY (account_id_from) REFERENCES account(id),
                             FOREIGN KEY (account_id_to) REFERENCES account(id)
);

CREATE TABLE balance (
                        id VARCHAR(36) PRIMARY KEY,
                        account_id VARCHAR(36),
                        balance DECIMAL(10, 2) NOT NULL,
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        FOREIGN KEY (account_id) REFERENCES account(id)
);

INSERT INTO client (id, name, email, created_at, updated_at)
VALUES
    ('client1', 'Client Name 1', 'client1@email.com', NOW(), NOW()),
    ('client2', 'Client Name 2', 'client2@email.com', NOW(), NOW()),
    ('client3', 'Client Name 3', 'client3@email.com', NOW(), NOW()),
    ('client4', 'Client Name 4', 'client4@email.com', NOW(), NOW()),
    ('client5', 'Client Name 5', 'client5@email.com', NOW(), NOW());

INSERT INTO account (id, client_id, balance, created_at, updated_at)
VALUES
    ('1', 'client1', 100.00, NOW(), NOW()),
    ('2', 'client2', 200.00, NOW(), NOW()),
    ('3', 'client3', 300.00, NOW(), NOW()),
    ('4', 'client4', 400.00, NOW(), NOW()),
    ('5', 'client5', 500.00, NOW(), NOW());


