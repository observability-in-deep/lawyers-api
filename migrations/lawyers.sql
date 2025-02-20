CREATE TABLE if not exists lawyers (
    id UUID PRIMARY KEY,
    name TEXT ,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    oab TEXT UNIQUE,
    email TEXT UNIQUE,
    phone TEXT UNIQUE
);


INSERT INTO lawyers (id, name, created_at, updated_at, deleted_at, oab, email, phone)
VALUES
(gen_random_uuid(), 'Jane Smith', '2021-02-01 00:00:00', '2021-02-01 00:00:00', NULL, '12345-abc-678', 'jane.smith@gmail.com', '88888888888'),
(gen_random_uuid(), 'Alice Johnson', '2021-03-01 00:00:00', '2021-03-01 00:00:00', NULL, '67890-def-123', 'alice.johnson@gmail.com', '77777777777'),
(gen_random_uuid(), 'Bob Brown', '2021-04-01 00:00:00', '2021-04-01 00:00:00', NULL, '23456-ghi-789', 'bob.brown@gmail.com', '66666666666'),
(gen_random_uuid(), 'Charlie Davis', '2021-05-01 00:00:00', '2021-05-01 00:00:00', NULL, '34567-jkl-012', 'charlie.davis@gmail.com', '55555555555'),
(gen_random_uuid(), 'Diana Evans', '2021-06-01 00:00:00', '2021-06-01 00:00:00', NULL, '45678-mno-345', 'diana.evans@gmail.com', '44444444444');
