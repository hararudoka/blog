CREATE TABLE customer (
    id SERIAL,
    username VARCHAR,
    password VARCHAR,
    role VARCHAR,
    is_admin BOOLEAN,
    created_at TIMESTAMP
);

CREATE TABLE post (
    id SERIAL,
    customer_id INT,
    title VARCHAR,
    content TEXT,
    created_at TIMESTAMP
);

CREATE TABLE comment (
    id SERIAL,
    post_id INT,
    customer_id INT,
    content TEXT,
    created_at TIMESTAMP
);

CREATE TABLE security (
    id SERIAL,
    token VARCHAR,
    customer_id INT
);