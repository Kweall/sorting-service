CREATE TABLE IF NOT EXISTS numbers (
    id serial PRIMARY KEY,
    value integer NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now()
);
