ALTER TABLE messages
ADD COLUMN type TEXT NOT NULL DEFAULT 'text',
ADD COLUMN shared_task_source_id UUID,
ADD COLUMN shared_task_title TEXT NOT NULL DEFAULT '',
ADD COLUMN shared_task_notes TEXT NOT NULL DEFAULT '',
ADD COLUMN shared_task_group_name TEXT NOT NULL DEFAULT '',
ADD COLUMN shared_task_group_color TEXT NOT NULL DEFAULT '';

CREATE INDEX idx_messages_type ON messages(type);
