CREATE TABLE IF NOT EXISTS "users"(
    "tg_id" INT PRIMARY KEY,
    "tg_name" VARCHAR(50) NOT NULL,
    "fullname" VARCHAR(100),
    "phone_number" VARCHAR(50) UNIQUE,
    "step" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
