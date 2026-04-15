DROP INDEX IF EXISTS idx_messages_type;

ALTER TABLE messages
DROP COLUMN IF EXISTS shared_task_group_color,
DROP COLUMN IF EXISTS shared_task_group_name,
DROP COLUMN IF EXISTS shared_task_notes,
DROP COLUMN IF EXISTS shared_task_title,
DROP COLUMN IF EXISTS shared_task_source_id,
DROP COLUMN IF EXISTS type;
