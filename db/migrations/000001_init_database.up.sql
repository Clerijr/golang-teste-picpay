-- Enable the uuid-ossp extension for generating UUIDs
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar(30) NOT NULL,
    lastname varchar(30),
    type varchar(8) CHECK (type IN ('fisica', 'juridica')) NOT NULL,
    document varchar(14) NOT NULL,
    email varchar(50) NOT NULL,
    password varchar(200) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,
    UNIQUE (email),
    UNIQUE (document)
);

CREATE TABLE IF NOT EXISTS wallets
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id uuid REFERENCES users (id),
    balance INTEGER
);

CREATE TABLE IF NOT EXISTS transactions
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    transfer_date TIMESTAMP NOT NULL DEFAULT now(),
    sender_wallet uuid REFERENCES wallets (id),
    receiver_wallet uuid REFERENCES wallets (id),
    amount INTEGER
);
