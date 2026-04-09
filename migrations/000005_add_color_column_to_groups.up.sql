ALTER TABLE groups ADD COLUMN color TEXT;

UPDATE groups
SET color = group_color
WHERE color IS NULL;

ALTER TABLE groups
ALTER COLUMN color SET NOT NULL;
