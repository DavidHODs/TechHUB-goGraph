BEGIN;

CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS pg_stat_statements;


CREATE SCHEMA IF NOT EXISTS tech;

CREATE TABLE IF NOT EXISTS tech.tags(
	id uuid NOT NULL DEFAULT gen_random_uuid () PRIMARY KEY UNIQUE,
	name VARCHAR (127) NOT NULL
);

END;