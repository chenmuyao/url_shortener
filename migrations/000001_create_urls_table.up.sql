CREATE SEQUENCE urls_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."urls" (
    "id" bigint DEFAULT nextval('urls_id_seq') NOT NULL,
    "url" character varying NOT NULL,
    CONSTRAINT "urls_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "urls_url" UNIQUE ("url")
) WITH (oids = false);
