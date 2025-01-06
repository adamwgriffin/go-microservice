CREATE TABLE "contacts" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "contacts" ("first_name");
CREATE INDEX ON "contacts" ("last_name");
