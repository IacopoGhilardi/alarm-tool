DO
$do$
BEGIN
   IF NOT EXISTS (
      SELECT
      FROM   pg_catalog.pg_database
      WHERE  datname = 'alarm_tool') THEN

      CREATE DATABASE alarm_tool;
   END IF;
END
$do$;

\c alarm_tool;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    uuid UUID UNIQUE DEFAULT gen_random_uuid() NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_uuid ON users(uuid);
