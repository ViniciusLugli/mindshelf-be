ALTER TABLE messages
ADD COLUMN IF NOT EXISTS imported_task_id UUID;

CREATE INDEX IF NOT EXISTS idx_messages_imported_task_id ON messages(imported_task_id);

ALTER TABLE messages
ADD CONSTRAINT fk_messages_imported_task
FOREIGN KEY (imported_task_id) REFERENCES tasks(id) ON DELETE SET NULL;
