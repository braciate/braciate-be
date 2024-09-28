CREATE TABLE assets (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) REFERENCES users(id) ON DELETE CASCADE,
    nomination_id VARCHAR(255) REFERENCES nominations(id) ON DELETE CASCADE,
    lkm_id VARCHAR(255) REFERENCES lkms(id) ON DELETE CASCADE,
    url VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
