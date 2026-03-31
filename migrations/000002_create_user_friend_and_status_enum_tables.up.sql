CREATE TYPE status AS ENUM ('pending', 'accepted', 'rejected');

CREATE TABLE user_friends (
    user_id   UUID REFERENCES users(id) ON DELETE CASCADE,
    friend_id UUID REFERENCES users(id) ON DELETE CASCADE,
    status    status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (user_id, friend_id)
);

CREATE INDEX idx_user_friends_friend_id ON user_friends(friend_id);