CREATE TABLE if not exists customers (
    id UUID PRIMARY KEY,
    name TEXT ,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    cpf TEXT UNIQUE,
    email TEXT UNIQUE,
    phone TEXT UNIQUE,
    folder TEXT,
    lawyer_id UUID,
    CONSTRAINT fk_lawyer
        FOREIGN KEY (lawyer_id)
        REFERENCES lawyers(id)
        ON DELETE SET NULL
);
