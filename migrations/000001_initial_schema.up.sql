

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION citext;

CREATE TYPE token_type AS ENUM ('confirmation', 'password_reset');
CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4 (),
    username VARCHAR(24) NOT NULL,
    email  citext UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    password text NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    verified BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);

CREATE TABLE tokens (
  user_id uuid NOT NULL REFERENCES users(id) DEFERRABLE INITIALLY DEFERRED,
  token uuid NOT NULL DEFAULT uuid_generate_v4 (),
  consumed boolean DEFAULT false,
  kind token_type NOT NULL, 
  expiration timestamp NOT NULL DEFAULT (NOW() + interval '2 day'),
  PRIMARY KEY (token)
)