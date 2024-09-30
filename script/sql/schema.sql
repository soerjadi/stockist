CREATE TYPE user_role AS ENUM('user','admin');

CREATE TABLE IF NOT EXISTS users (
    id          BIGSERIAL PRIMARY KEY NOT NULL,
    name        varchar(100) NOT NULL,
    email       varchar(155) NOT NULL,
    phone_number varchar(13) NOT NULL,
    address     varchar(255) NOT NULL,
    role        user_role NOT NULL,
    password    varchar(255) NOT NULL,
    salt        varchar(255) NOT NULL,
    created_at  timestamp default current_timestamp
);

CREATE TABLE IF NOT EXISTS stores (
    id          BIGSERIAL PRIMARY KEY NOT NULL,
    name        varchar(100) NOT NULL,
    description varchar(255) NOT NULL,
    address     varchar(100) NOT NULL,
    manager_id  BIGINT NOT NULL,
    created_at  timestamp default current_timestamp,
    CONSTRAINT fk_stores_users
        FOREIGN KEY (manager_id)
        REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS products (
    id              BIGSERIAL PRIMARY KEY NOT NULL,
    name            varchar(100) NOT NULL,
    description     varchar(255) NOT NULL,
    weight          int NOT NULL,
    price           BIGINT NOT NULL,
    store_id        BIGINT NOT NULL,
    stock           BIGINT NOT NULL,
    images          varchar(255),
    created_at      timestamp default current_timestamp,
    updated_at      timestamp,
    CONSTRAINT fk_products_stores
        FOREIGN KEY (store_id)
        REFERENCES stores (id)
);

CREATE TYPE order_status AS ENUM('paid', 'unpaid', 'created', 'stale', 'delivery');

CREATE TABLE IF NOT EXISTS orders (
    id              BIGSERIAL PRIMARY KEY NOT NULL,
    user_id         BIGINT NOT NULL,
    store_id        BIGINT NOT NULL,
    total_price     BIGINT NOT NULL,
    total_amount    BIGINT NOT NULL,
    status          order_status NOT NULL,
    created_at      timestamp default current_timestamp,
    updated_at      timestamp,
    CONSTRAINT fk_orders_users
        FOREIGN KEY (user_id)
        REFERENCES users (id),
    CONSTRAINT fk_orders_stores
        FOREIGN KEY (store_id)
        REFERENCES stores (id)
);

CREATE TABLE IF NOT EXISTS order_items (
    id          BIGSERIAL PRIMARY KEY NOT NULL,
    order_id    BIGINT NOT NULL,
    product_id  BIGINT NOT NULL,
    amount      BIGINT NOT NULL,
    price       BIGINT NOT NULL,
    created_at  timestamp default current_timestamp,
    CONSTRAINT fk_order_items_orders
        FOREIGN KEY (order_id)
        REFERENCES orders (id),
    CONSTRAINT fk_order_items_products
        FOREIGN KEY (product_id)
        REFERENCES products (id)
);

CREATE TYPE warehouse_status AS ENUM('active', 'inactive');

CREATE TABLE IF NOT EXISTS warehouses (
    id          BIGSERIAL PRIMARY KEY NOT NULL,
    name        varchar(100) NOT NULL,
    address     varchar(155) NOT NULL,
    status      warehouse_status NOT NULL,
    quota       BIGINT NOT NULL,
    stock       BIGINT NOT NULL,
    created_at  timestamp default current_timestamp,
    updated_at  timestamp,
    status_updated_at timestamp
);

CREATE TABLE IF NOT EXISTS warehouse_products (
    id              BIGSERIAL PRIMARY KEY NOT NULL,
    warehouse_id    BIGINT NOT NULL,
    product_id      BIGINT NOT NULL,
    stock           BIGINT NOT NULL,
    created_at      timestamp default current_timestamp,
    updated_at      timestamp,
    CONSTRAINT fk_warehouse_products_warehouses
        FOREIGN KEY (warehouse_id)
        REFERENCES warehouses (id),
    CONSTRAINT fx_warehouse_products_products
        FOREIGN KEY (product_id)
        REFERENCES products (id)
);

CREATE TYPE warehouse_trf_status AS ENUM('received','inprogress','process_transfer','transfered');

CREATE TABLE IF NOT EXISTS warehouse_trf_logs (
    id              BIGSERIAL PRIMARY KEY NOT NULL,
    target_id       BIGINT NOT NULL,
    target_class    varchar(55) NOT NULL,
    product_id      BIGINT NOT NULL,
    amount          BIGINT NOT NULL,
    note            varchar(255) NOT NULL,
    status          warehouse_trf_status NOT NULL,
    created_at      timestamp default current_timestamp,
    updated_at      timestamp,
    CONSTRAINT fk_warehouse_trf_logs_products
        FOREIGN KEY (product_id)
        REFERENCES products (id)
);