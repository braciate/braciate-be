CREATE TABLE nominations (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    categories_id VARCHAR(255) REFERENCES categories(id)
);
