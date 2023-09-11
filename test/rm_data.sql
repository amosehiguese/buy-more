-- A handy script to remove all data from the database but still preserves the table structure
-- cd test/
-- run -> psql <db_name> <user_name>
-- input password
-- \i rm_data.sql -> to remove data
DO $$
DECLARE
    table_name text;
BEGIN
    FOR table_name IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public')
    LOOP
        EXECUTE 'DELETE FROM ' || table_name;
    END LOOP;
END $$;
