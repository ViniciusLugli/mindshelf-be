CREATE INDEX idx_messages_receiver_sender_read_created
ON messages(receiver_id, sender_id, read_at, created_at);
