CREATE TYPE users_status_enum AS ENUM ('active', 'inactive', 'blocked', 'other');
CREATE TYPE users_role_enum AS ENUM ('admin', 'merchant', 'customer');

CREATE TABLE IF NOT EXISTS users
(
    -- column
    id            UUID PRIMARY KEY NOT NULL,
    username      VARCHAR(25) UNIQUE,
    password      VARCHAR(100),
    email         VARCHAR(50) UNIQUE,
    phone_number  VARCHAR(15) UNIQUE,
    user_status   users_status_enum NOT NULL DEFAULT 'active',
    role          users_role_enum NOT NULL,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ
    );

CREATE TABLE IF NOT EXISTS wallets
(
    -- column
    id            UUID PRIMARY KEY NOT NULL,
    user_id       UUID REFERENCES users (id),
    rekening_user VARCHAR(100) NOT NULL,
    balance       BIGINT,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ
    );

CREATE TABLE IF NOT EXISTS categories
(
    -- column
    id            UUID PRIMARY KEY NOT NULL,
    type          VARCHAR(100) NOT NULL,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ,
    deleted_at    TIMESTAMPTZ
    );

CREATE TABLE IF NOT EXISTS products
(
    -- column
    id            UUID PRIMARY KEY NOT NULL,
    user_id       UUID REFERENCES users (id),
    product_name  VARCHAR(100) NOT NULL,
    price         BIGINT NOT NULL,
    stock         BIGINT NOT NULL,
    category_id   UUID REFERENCES categories (id),
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ,
    deleted_at    TIMESTAMPTZ
    );

CREATE TABLE IF NOT EXISTS carts
(
    -- column
    id            UUID PRIMARY KEY NOT NULL,
    user_id       UUID REFERENCES users (id),
    product_id    UUID REFERENCES products (id),
    quantity      BIGINT NOT NULL,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ
    );

CREATE TABLE IF NOT EXISTS transactions
(
    -- column
    id            UUID PRIMARY KEY NOT NULL,
    user_id       UUID REFERENCES users (id),
    product_id    UUID REFERENCES products (id),
    quantity      BIGINT NOT NULL,
    total_price   BIGINT NOT NULL,
    date          DATE
    );
