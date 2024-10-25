-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS public.links;

CREATE TABLE public.links (
  "id" bigserial primary key,
  "short" varchar not null,
  "original_url" varchar not null,
  "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS u_links_short_idx ON public.links (short);
CREATE UNIQUE INDEX IF NOT EXISTS u_links_original_url_idx ON public.links (original_url);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';

-- +goose StatementEnd