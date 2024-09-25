CREATE TABLE lkms (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id VARCHAR(255) REFERENCES categories(id) ON DELETE CASCADE,
    logo_link VARCHAR(255),
    type SMALLINT NOT NULL
);