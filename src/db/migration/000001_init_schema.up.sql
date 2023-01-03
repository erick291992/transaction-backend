CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "document_number" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "operations_types" (
  "id" bigserial PRIMARY KEY,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);