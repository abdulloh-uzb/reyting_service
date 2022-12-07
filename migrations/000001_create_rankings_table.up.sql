CREATE TABLE IF NOT EXISTS rankings (
    name varchar NOT NULL,
    ranking int NOT NULL,
    description varchar NOT NULL, 
    post_id int NOT NULL,
    customer_id int NOT NULL,
    created_at timestamptz NULL DEFAULT now(),
    deleted_at timestamptz NULL,
    updated_at timestamptz NULL
)