/*
 Navicat Premium Data Transfer

 Source Server         : local_domainadmin
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 30/12/2021 15:08:38
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for domains
-- ----------------------------
DROP TABLE IF EXISTS "domains";
CREATE TABLE "domains" (
  "id" integer,
  "created_at" datetime,
  "updated_at" datetime,
  "deleted_at" datetime,
  "name" text,
  PRIMARY KEY ("id"),
  UNIQUE ("name" ASC)
);

-- ----------------------------
-- Table structure for ip_addrs
-- ----------------------------
DROP TABLE IF EXISTS "ip_addrs";
CREATE TABLE "ip_addrs" (
  "id" integer,
  "created_at" datetime,
  "updated_at" datetime,
  "deleted_at" datetime,
  "ip" text,
  "domain_id" integer,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_domains_ip_addrs" FOREIGN KEY ("domain_id") REFERENCES "domains" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);

-- ----------------------------
-- Indexes structure for table domains
-- ----------------------------
CREATE INDEX "idx_domains_deleted_at"
ON "domains" (
  "deleted_at" ASC
);

-- ----------------------------
-- Indexes structure for table ip_addrs
-- ----------------------------
CREATE INDEX "idx_ip_addrs_deleted_at"
ON "ip_addrs" (
  "deleted_at" ASC
);

PRAGMA foreign_keys = true;
