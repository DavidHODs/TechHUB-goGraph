CREATE TABLE IF NOT EXISTS users (
	id uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
	name VARCHAR (127) NOT NULL,
	email VARCHAR (127) NOT NULL,
	password VARCHAR (127) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
);
