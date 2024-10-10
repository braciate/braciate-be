CREATE TABLE user_votes (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) REFERENCES users(id) ON DELETE CASCADE,
    nomination_id VARCHAR(255) REFERENCES nominations(id) ON DELETE CASCADE,
    lkm_id VARCHAR(255) REFERENCES lkms(id) ON DELETE CASCADE,
    CONSTRAINT user_lkm_nomination_unique UNIQUE (user_id, lkm_id, nomination_id)
);