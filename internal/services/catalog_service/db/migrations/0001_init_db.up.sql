CREATE TABLE IF NOT EXISTS types
(
    id         serial primary key,
    type       text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS brands
(
    id         serial primary key,
    brand      text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS items
(
    id                  serial primary key,
    name                text,
    description         text,
    price               numeric,
    picture_file_name   text,
    picture_uri         text,
    type_id             integer references types (id),
    brand_id            integer references brands (id),
    available_stock     integer,
    restock_threshold   integer,
    max_stock_threshold integer,
    on_reorder          boolean,
    created_at          timestamp with time zone,
    updated_at          timestamp with time zone
);
