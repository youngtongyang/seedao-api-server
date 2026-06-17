-- Run as postgres superuser on first install (native PostgreSQL)
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'seedao') THEN
        CREATE ROLE seedao WITH LOGIN PASSWORD 'seedao';
    END IF;
END
$$;

SELECT 'CREATE DATABASE seedao_api OWNER seedao'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'seedao_api')\gexec

GRANT ALL PRIVILEGES ON DATABASE seedao_api TO seedao;
