-- +goose Up
-- Create "companies" table
CREATE TABLE "companies" (
  "id" uuid NOT NULL,
  "name" text,
  "location" text,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);
-- Create "users" table
CREATE TABLE "users" (
  "id" uuid NOT NULL,
  "username" text,
  "email" text,
  "password" text,
  "name" text,
  "company_id" uuid,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_company" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "users_username" to table: "users"
CREATE UNIQUE INDEX "users_username" ON "users" ("username");
-- Create index "users_email" to table: "users"
CREATE UNIQUE INDEX "users_email" ON "users" ("email");
-- Create "test_runs" table
CREATE TABLE "test_runs" (
  "id" uuid NOT NULL,
  "company_id" uuid,
  "started_at" timestamp,
  "finished_at" timestamp,
  "status" text,
  "total_tests" integer,
  "passed_tests" integer,
  "failed_tests" integer,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_test_runs_company" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "test_suites" table
CREATE TABLE "test_suites" (
  "id" uuid NOT NULL,
  "company_id" uuid,
  "name" text,
  "method" text,
  "url" text,
  "headers" text,
  "expected_status" integer,
  "expected_body" text,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_test_suites_company" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "test_results" table
CREATE TABLE "test_results" (
  "id" uuid NOT NULL,
  "test_run_id" uuid,
  "endpoint_test_id" uuid,
  "status" text,
  "response_status" integer,
  "response_body" text,
  "response_time_ms" integer,
  "error_message" text,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_test_results_endpoint_test" FOREIGN KEY ("endpoint_test_id") REFERENCES "test_suites" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_test_results_test_run" FOREIGN KEY ("test_run_id") REFERENCES "test_runs" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- +goose Down
DROP TABLE IF EXISTS "test_results";
DROP TABLE IF EXISTS "test_suites";
DROP TABLE IF EXISTS "test_runs";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "companies";
