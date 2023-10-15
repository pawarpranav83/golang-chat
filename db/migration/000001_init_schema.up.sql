CREATE TABLE
    "rooms" (
        "id" bigserial PRIMARY KEY,
        "roomname" varchar UNIQUE NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now ())
    );

CREATE TABLE
    "users" (
        "id" bigserial PRIMARY KEY,
        "username" varchar UNIQUE NOT NULL,
        "role" varchar NOT NULL DEFAULT 'user',
        "email" varchar UNIQUE NOT NULL,
        "password" varchar NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now ())
    );

CREATE TABLE
    "userroom" (
        "room_id" bigint NOT NULL,
        "user_id" bigint NOT NULL,
        PRIMARY KEY ("room_id", "user_id")
    );

CREATE INDEX ON "rooms" ("roomname");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "userroom" ("room_id");

CREATE INDEX ON "userroom" ("user_id");

ALTER TABLE "userroom" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "userroom" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");