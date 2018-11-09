DROP TABLE IF EXISTS Users CASCADE ;
DROP TABLE IF EXISTS forums CASCADE ;
DROP TABLE IF EXISTS threads CASCADE;

CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS Users (
  about    TEXT,
  email    CITEXT COLLATE "ucs_basic" NOT NULL UNIQUE,
  fullname VARCHAR,
  nickname CITEXT COLLATE "ucs_basic" NOT NULL UNIQUE
);

CREATE TABLE forums (
  posts INTEGER DEFAULT 0,
	slug CITEXT PRIMARY KEY,
	threads INTEGER DEFAULT 0,
	title CITEXT,
	author CITEXT REFERENCES users(nickname)
);


CREATE TABLE threads(
  author CITEXT REFERENCES users(nickname),
  created TIMESTAMP WITH TIME ZONE,
  forum CITEXT REFERENCES forums(slug),
	id BIGSERIAL PRIMARY KEY,
	message TEXT NOT NULL,
	slug CITEXT UNIQUE,
	title TEXT NOT NULL,
	votes BIGINT DEFAULT 0
);

CREATE OR REPLACE FUNCTION thread_inc() RETURNS trigger AS $$
BEGIN
  UPDATE forums SET threads = threads + 1
  WHERE slug = NEW.forum;
  RETURN NEW;
END;
$$ language plpgsql;

DROP TRIGGER IF EXISTS thread_inc ON Threads;

CREATE TRIGGER thread_inc AFTER INSERT ON Threads
  FOR EACH ROW EXECUTE PROCEDURE thread_inc();