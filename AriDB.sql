CREATE TABLE "users" (
  "id" int PRIMARY KEY,
  "user_id" int,
  "email" varchar UNIQUE,
  "project_included" varchar
);

CREATE TABLE "projects" (
  "project_id" int PRIMARY KEY,
  "project_name" varchar,
  "created_at" varchar,
  "erp_standards" varchar,
  "files" varchar
);

CREATE TABLE "csv_files" (
  "csv_id" int PRIMARY KEY,
  "unconverted_csv" varchar,
  "converted_csv" varchar,
  "comparrison_csv" varchar
);

COMMENT ON COLUMN "projects"."created_at" IS 'When project created';

COMMENT ON COLUMN "projects"."erp_standards" IS 'Headers of comparrison_csv';

ALTER TABLE "projects" ADD FOREIGN KEY ("project_id") REFERENCES "users" ("project_included");

ALTER TABLE "projects" ADD FOREIGN KEY ("project_name") REFERENCES "users" ("project_included");

ALTER TABLE "csv_files" ADD FOREIGN KEY ("comparrison_csv") REFERENCES "projects" ("erp_standards");

ALTER TABLE "csv_files" ADD FOREIGN KEY ("unconverted_csv") REFERENCES "projects" ("files");

ALTER TABLE "csv_files" ADD FOREIGN KEY ("converted_csv") REFERENCES "projects" ("files");
