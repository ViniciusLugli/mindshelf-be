DROP INDEX IF EXISTS idx_messages_deleted_at;

ALTER TABLE messages
DROP COLUMN IF EXISTS deleted_at,
DROP COLUMN IF EXISTS updated_at;
