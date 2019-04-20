\c spajamdojo2019
-- Table: public.messages

-- DROP TABLE public.messages;

CREATE TABLE public.terms
(
    term_id serial NOT NULL,
    t_from  character varying(256),
    t_to  character varying(256),
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.scores
    OWNER to root;
