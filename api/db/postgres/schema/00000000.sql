DROP EXTENSION IF EXISTS pgcrypto;
DROP TABLE IF EXISTS sellers;
DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS inventory;
DROP TABLE IF EXISTS products;


CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE
    IF NOT EXISTS products
(
    id            UUID PRIMARY KEY        DEFAULT gen_random_uuid(),
    name          VARCHAR(200)   NOT NULL,
    released_date DATE           NOT NULL,
    model         VARCHAR(100),
    price         NUMERIC(12, 2) NOT NULL,
    manufacturer  VARCHAR(200)   NOT NULL,
    is_active     BOOLEAN                 DEFAULT TRUE,
    created_by    VARCHAR(200)   NOT NULL,
    created_at    TIMESTAMPTZ    NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_by    VARCHAR(200)   NOT NULL,
    updated_at    TIMESTAMPTZ    NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);

CREATE TABLE
    IF NOT EXISTS inventory
(
    id         UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    product_id UUID REFERENCES products (id) ON DELETE CASCADE,
    quantity   INTEGER      NOT NULL CHECK (quantity > 0),
    created_by VARCHAR(200) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_by VARCHAR(200) NOT NULL,
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);

CREATE TABLE
    IF NOT EXISTS users
(
    id         UUID PRIMARY KEY             DEFAULT gen_random_uuid(),
    username   VARCHAR(50) UNIQUE  NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    is_active  BOOLEAN                      DEFAULT TRUE,
    created_by VARCHAR(200)        NOT NULL,
    created_at TIMESTAMPTZ         NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_by VARCHAR(200)        NOT NULL,
    updated_at TIMESTAMPTZ         NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);

CREATE TABLE
    IF NOT EXISTS orders
(
    id          UUID PRIMARY KEY        DEFAULT gen_random_uuid(),
    user_id     UUID REFERENCES users (id),
    total_items INTEGER        NOT NULL CHECK (total_items > 0),
    total_price NUMERIC(12, 2) NOT NULL,
    status      VARCHAR(20) CHECK (
        status IN ('PENDING', 'SHIPPED', 'DELIVERED', 'CANCELLED')
        )                               DEFAULT 'PENDING',
    created_at  TIMESTAMPTZ    NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);

CREATE TABLE
    IF NOT EXISTS carts
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID REFERENCES users (id),
    total_items INTEGER        NOT NULL CHECK (total_items > 0),
    total_price NUMERIC(12, 2) NOT NULL
);

CREATE TABLE
    IF NOT EXISTS order_items
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id   UUID REFERENCES orders (id) ON DELETE CASCADE,
    product_id UUID REFERENCES products (id),
    quantity   INTEGER        NOT NULL CHECK (quantity > 0),
    price      NUMERIC(12, 2) NOT NULL,
    UNIQUE (order_id, product_id)
);

CREATE TABLE
    IF NOT EXISTS cart_items
(
    id         UUID PRIMARY KEY        DEFAULT gen_random_uuid(),
    cart_id    UUID REFERENCES carts (id) ON DELETE CASCADE,
    product_id UUID REFERENCES products (id),
    quantity   INTEGER        NOT NULL CHECK (quantity > 0),
    price      NUMERIC(12, 2) NOT NULL,
    created_at TIMESTAMPTZ    NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_at TIMESTAMPTZ    NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    UNIQUE (cart_id, product_id)
);

CREATE TABLE
    IF NOT EXISTS sellers
(
    id         UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    user_id    UUID REFERENCES users (id),
    is_active  BOOLEAN               DEFAULT TRUE,
    created_by VARCHAR(200) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_by VARCHAR(200) NOT NULL,
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);
