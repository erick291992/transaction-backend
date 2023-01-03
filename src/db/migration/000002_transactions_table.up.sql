CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "operation_type" bigint NOT NULL,
  "amount" decimal NOT NULL DEFAULT 0,
  "event_date" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "transactions" ("account_id");

CREATE INDEX ON "transactions" ("operation_type");

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("operation_type") REFERENCES "operations_types" ("id");
