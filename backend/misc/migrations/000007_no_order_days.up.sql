ALTER TABLE users
ADD COLUMN IF NOT EXISTS no_order_days text[] DEFAULT '{}';

-- Ensure existing rows are updated to have the default value
UPDATE users
SET no_order_days = '{}'
WHERE no_order_days IS NULL;
