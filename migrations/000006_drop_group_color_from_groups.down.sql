ALTER TABLE groups ADD COLUMN group_color TEXT;

UPDATE groups
SET group_color = color
WHERE group_color IS NULL;

ALTER TABLE groups
ALTER COLUMN group_color SET NOT NULL;
