\c cown
-- Table: public.messages

-- DROP TABLE public.messages;

CREATE TABLE public.bets
(
    bet_id serial NOT NULL,
    user_id integer,
    amount integer,
    team_id integer,
    term integer,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "user_id_FK" FOREIGN KEY (user_id)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.scores
    OWNER to root;
