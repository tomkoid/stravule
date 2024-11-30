ALTER TABLE filters
ADD COLUMN IF NOT EXISTS weight integer NOT NULL DEFAULT 1;

-- Ensure existing rows are updated to have the default value
UPDATE filters
SET weight = 1
WHERE category = 'include';

UPDATE filters
SET weight = 2
WHERE category = 'exclude';
