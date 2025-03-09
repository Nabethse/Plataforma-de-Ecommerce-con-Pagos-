CREATE TABLE payments (
    id UUID PRIMARY KEY,
    order_id UUID REFERENCES orders(id),
    amount DECIMAL(10,2),
    status VARCHAR(50)
);