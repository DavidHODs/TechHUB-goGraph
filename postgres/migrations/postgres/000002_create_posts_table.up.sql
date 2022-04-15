BEGIN;

CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS pg_stat_statements;


CREATE SCHEMA IF NOT EXISTS tech;

-- CREATE OR REPLACE FUNCTION set_updated_at() RETURNS trigger AS $$
-- BEGIN
--   IF (
--     NEW IS DISTINCT FROM OLD AND
--     NEW.updated_at IS NOT DISTINCT FROM OLD.updated_at
--   ) THEN
--     NEW.updated_at := current_timestamp;
--   END IF;
--   RETURN NEW;
-- END;

CREATE TABLE IF NOT EXISTS tech.posts (
	id uuid NOT NULL DEFAULT gen_random_uuid () PRIMARY KEY,
	body VARCHAR (1027) NOT NULL,
	shared_body VARCHAR (1027),
	image VARCHAR (127),
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
	update_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
	shared_at TIMESTAMP WITH TIME ZONE,
	author uuid NOT NULL REFERENCES tech.users (id) ON DELETE CASCADE ON UPDATE CASCADE,
	shared_user uuid REFERENCES tech.users (id) ON DELETE CASCADE ON UPDATE CASCADE,
	likes uuid REFERENCES tech.users (id) ON DELETE CASCADE ON UPDATE CASCADE,
	Dislikes uuid REFERENCES tech.users (id) ON DELETE CASCADE ON UPDATE CASCADE,
	Tags uuid REFERENCES tech.tags (id) ON DELETE CASCADE ON UPDATE CASCADE
);

END;