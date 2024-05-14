CREATE TABLE "user" (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    username varchar(150),
    password varchar(128),
    first_name varchar(150),
    last_name varchar(150),
    email varchar(254),
    phone varchar(50),
    is_superuser bool,
    is_staff bool,
    is_active bool,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW()
)