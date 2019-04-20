\c spajamdojo2019
-- Table: public.users

-- DROP TABLE public.users;

CREATE TABLE public.users
(
    user_id serial NOT NULL,
    user_name character varying(256),
    avater_path character varying(256),
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "Users_pkey" PRIMARY KEY (user_id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE spajamdojo2019.public.users
    OWNER to root;
