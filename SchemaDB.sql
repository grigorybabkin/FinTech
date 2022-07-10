CREATE TABLE "urls" (
  "full_url" varchar UNIQUE NOT NULL,
  "short_url" varchar UNIQUE NOT NULL
);

CREATE INDEX "both_urls" ON "urls" ("full_url", "short_url");

CREATE UNIQUE INDEX ON "urls" ("full_url");

CREATE UNIQUE INDEX ON "urls" ("short_url");
