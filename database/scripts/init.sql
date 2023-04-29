-- Database: test_db

-- DROP DATABASE IF EXISTS test_db;

CREATE DATABASE test_db
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

ALTER DATABASE test_db OWNER TO postgres;

\connect test_db

-- Table: public.users

-- DROP TABLE IF EXISTS public.users;

CREATE TABLE IF NOT EXISTS public.users
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    refresh_token character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT unique_users_email UNIQUE (email)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;
-- Index: users_id_index

-- DROP INDEX IF EXISTS public.users_id_index;

CREATE UNIQUE INDEX IF NOT EXISTS users_id_index
    ON public.users USING btree
    (id ASC NULLS LAST)
    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    CLUSTER ON users_id_index;