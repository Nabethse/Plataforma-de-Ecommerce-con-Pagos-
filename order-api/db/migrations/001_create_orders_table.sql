CREATE TABLE orders (
    id UUID PRIMARY KEY,
    amount DECIMAL(10,2),
    status VARCHAR(50)
);