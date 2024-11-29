ALTER TABLE users
ADD COLUMN IF NOT EXISTS order_days_exceptions integer[] DEFAULT '{}';

-- Ensure existing rows are updated to have the default value
UPDATE users
SET order_days_exceptions = '{}' 
WHERE order_days_exceptions IS NULL;
