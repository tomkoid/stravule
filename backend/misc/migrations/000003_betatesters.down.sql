-- Remove the is_beta_tester column
DO $$
BEGIN
    -- Check if the column exists before attempting to drop it
    IF EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'users' AND column_name = 'is_beta_tester'
    ) THEN
        ALTER TABLE users
        DROP COLUMN is_beta_tester;
    END IF;
END $$;
