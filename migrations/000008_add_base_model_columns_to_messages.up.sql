ALTER TABLE messages
ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ DEFAULT NOW(),
ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ;

UPDATE messages
SET updated_at = created_at
WHERE updated_at IS NULL;

ALTER TABLE messages
ALTER COLUMN updated_at SET NOT NULL;

CREATE INDEX IF NOT EXISTS idx_messages_deleted_at ON messages(deleted_at);
