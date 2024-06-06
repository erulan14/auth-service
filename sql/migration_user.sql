CREATE TABLE "user" (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    username varchar(150) unique,
    password varchar(128),
    first_name varchar(150) NOT NULL DEFAULT '',
    last_name varchar(150) NOT NULL DEFAULT '',
    email varchar(254) NOT NULL DEFAULT '',
    phone varchar(50) NOT NULL DEFAULT '',
    is_superuser bool DEFAULT FALSE,
    is_staff bool NOT NULL DEFAULT FALSE,
    is_active bool NOT NULL DEFAULT FALSE,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT current_timestamp,
    is_deleted bool NOT NULL DEFAULT FALSE
)


