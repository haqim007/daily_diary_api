/*
 Navicat Premium Data Transfer

 Source Server         : postgre local
 Source Server Type    : PostgreSQL
 Source Server Version : 90608
 Source Host           : localhost:5432
 Source Catalog        : diary
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 90608
 File Encoding         : 65001

 Date: 16/06/2020 02:30:12
*/


-- ----------------------------
-- Sequence structure for diaries_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."diaries_id_seq";
CREATE SEQUENCE "public"."diaries_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_sessions_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_sessions_id_seq";
CREATE SEQUENCE "public"."user_sessions_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."users_id_seq";
CREATE SEQUENCE "public"."users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for diaries
-- ----------------------------
DROP TABLE IF EXISTS "public"."diaries";
CREATE TABLE "public"."diaries" (
  "id" int4 NOT NULL DEFAULT nextval('diaries_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "user_id" int4,
  "date" timestamptz(6),
  "content" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Table structure for user_sessions
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_sessions";
CREATE TABLE "public"."user_sessions" (
  "id" int4 NOT NULL DEFAULT nextval('user_sessions_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "user_id" int4,
  "token" text COLLATE "pg_catalog"."default",
  "refresh_token" text COLLATE "pg_catalog"."default",
  "issued_at" timestamptz(6)
)
;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "fullname" text COLLATE "pg_catalog"."default",
  "birthday" timestamptz(6),
  "email" text COLLATE "pg_catalog"."default",
  "username" text COLLATE "pg_catalog"."default",
  "password" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."diaries_id_seq"
OWNED BY "public"."diaries"."id";
SELECT setval('"public"."diaries_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_sessions_id_seq"
OWNED BY "public"."user_sessions"."id";
SELECT setval('"public"."user_sessions_id_seq"', 2, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."users_id_seq"
OWNED BY "public"."users"."id";
SELECT setval('"public"."users_id_seq"', 2, true);

-- ----------------------------
-- Indexes structure for table diaries
-- ----------------------------
CREATE INDEX "idx_diaries_deleted_at" ON "public"."diaries" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table diaries
-- ----------------------------
ALTER TABLE "public"."diaries" ADD CONSTRAINT "diaries_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table user_sessions
-- ----------------------------
CREATE INDEX "idx_user_sessions_deleted_at" ON "public"."user_sessions" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table user_sessions
-- ----------------------------
ALTER TABLE "public"."user_sessions" ADD CONSTRAINT "user_sessions_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table users
-- ----------------------------
CREATE INDEX "idx_users_deleted_at" ON "public"."users" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table diaries
-- ----------------------------
ALTER TABLE "public"."diaries" ADD CONSTRAINT "user_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
