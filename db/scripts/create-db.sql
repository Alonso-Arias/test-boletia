-- Crear una nueva base de datos
CREATE DATABASE test_boletia_db;

-- Conectarse a la nueva base de datos
\c test_boletia_db;

-- Crear las tablas dentro del nuevo esquema
CREATE TABLE IF NOT EXISTS public.CURRENCIES (
    ID SERIAL PRIMARY KEY NOT NULL,
    CURRENCY TEXT NOT NULL,
    VALUE NUMERIC NOT NULL,
    TIMESTAMP TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS public.CALLS_LOG (
    ID SERIAL PRIMARY KEY NOT NULL,
    CALL_TIMESTAMP TIMESTAMP NOT NULL,
    RESPONSE_TIME_MS INT NOT NULL,
    STATUS TEXT NOT NULL
);
