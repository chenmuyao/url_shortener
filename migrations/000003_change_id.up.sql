ALTER TABLE urls ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS urls_id_seq;
