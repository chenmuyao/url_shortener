CREATE TABLE public.urls (
    id bigint NOT NULL,
    full_url varchar NOT NULL,
    CONSTRAINT urls_pk PRIMARY KEY (id),
    CONSTRAINT urls_unique UNIQUE (full_url)
);
