CREATE TABLE "categories" (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    is_active INT NOT NULL DEFAULT 1,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);