ALTER TABLE messages
DROP CONSTRAINT IF EXISTS fk_messages_imported_task;

DROP INDEX IF EXISTS idx_messages_imported_task_id;

ALTER TABLE messages
DROP COLUMN IF EXISTS imported_task_id;
