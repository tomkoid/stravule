-- Add the is_beta_tester column with a default value
DO $$
BEGIN
    -- Check if the column does not exist
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'users' AND column_name = 'is_beta_tester'
    ) THEN
        ALTER TABLE users
        ADD COLUMN is_beta_tester BOOLEAN NOT NULL DEFAULT FALSE;
    END IF;

    -- Ensure existing rows are updated to have the default value
    UPDATE users
    SET is_beta_tester = FALSE
    WHERE is_beta_tester IS NULL;
END $$;
