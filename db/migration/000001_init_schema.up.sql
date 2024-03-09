CREATE TABLE "users" (
    "id" uuid PRIMARY KEY, "name" varchar NOT NULL, "email" varchar UNIQUE NOT NULL, "password" varchar NOT NULL
);