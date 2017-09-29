--SELECT * FROM showtables()

CREATE TYPE showtable AS ("Schema" name,"Table" name);

CREATE OR REPLACE FUNCTION showtables() 
RETURNS SETOF showtable 
as $$
declare 
    t text :=E'\'r\'';
    blank text :=E'\'\'';
    pgc text :=E'\'pg_catalog\'';
    ins text :=E'\'information_schema\'';
    pgt text :=E'\'^pg_toast\'';
    query_tables text; 
begin 
    query_tables:='SELECT n.nspname as "Schema",
  c.relname as "Name"
FROM pg_catalog.pg_class c
     LEFT JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
WHERE c.relkind IN ('||t||','||blank||')
      AND n.nspname <> '||pgc||'
      AND n.nspname <> '||ins||'
      AND n.nspname !~ '||pgt||'
  AND pg_catalog.pg_table_is_visible(c.oid)
ORDER BY 1,2';
    return query EXECUTE query_tables;  
end; 
$$ LANGUAGE plpgsql;
DROP FUNCTION tablenames();
DROP TYPE tablename;
CREATE TYPE tablename AS ("Name" name); 
CREATE OR REPLACE FUNCTION tablenames() 
RETURNS SETOF tablename 
as $$
BEGIN
   return query EXECUTE 'SELECT "Table"  from showtables()';
END; 
$$ LANGUAGE plpgsql;