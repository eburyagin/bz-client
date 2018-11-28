DROP TABLE IF EXISTS public.clients;
CREATE TABLE public.clients
(
    id CHARACTER VARYING(25) PRIMARY KEY,
    code CHARACTER VARYING(20),
    name CHARACTER VARYING(100),
    note CHARACTER VARYING,
    status CHARACTER VARYING(20) NOT NULL,
    create_time TIMESTAMP DEFAULT current_timestamp
) WITH (
    OIDS = FALSE
);
