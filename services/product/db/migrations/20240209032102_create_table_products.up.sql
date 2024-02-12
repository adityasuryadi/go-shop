CREATE TABLE "products" (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    price float8,
    stock int,
    description TEXT,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);