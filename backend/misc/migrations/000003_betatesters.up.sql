ALTER TABLE users
ADD COLUMN IF NOT EXISTS is_beta_tester BOOLEAN NOT NULL DEFAULT FALSE;

-- Ensure existing rows are updated to have the default value
UPDATE users
SET is_beta_tester = FALSE
WHERE is_beta_tester IS NULL;
