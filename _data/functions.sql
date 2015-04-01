CREATE FUNCTION update_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$ BEGIN
     NEW.updated_at = CURRENT_TIMESTAMP;
     RETURN NEW;
  END;
  $$;


ALTER FUNCTION public.update_updated_at() OWNER TO {{DB_USER}};