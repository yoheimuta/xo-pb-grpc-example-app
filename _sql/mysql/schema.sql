CREATE DATABASE IF NOT EXISTS `test-xo-db`;
USE `test-xo-db`

-- users stores users' accounts.
CREATE TABLE users (
  -- user_id represents an user's id. This is UUID v4.
  user_id varchar(36) NOT NULL PRIMARY KEY,
  -- create_at represents a creation timestamp.
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  -- update_at represents a last update timestamp.
  updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE Now()
) ENGINE=InnoDB;

-- user_auths stores users' authentication data.
CREATE TABLE user_auths (
  -- user_id represents an id from the users table.
  user_id varchar(36) NOT NULL PRIMARY KEY,
  -- email represents an email address.
  email varchar(255) NOT NULL,
  -- password_hash represents a one-way hash of password.
  password_hash varchar(255) NOT NULL,
  -- create_at represents a creation timestamp.
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  -- update_at represents a last update timestamp.
  updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE Now(),
  CONSTRAINT FOREIGN KEY (user_id) REFERENCES users(user_id)
) ENGINE=InnoDB;

-- user_auths_email_unique_idx represents a unique constraint for user's email.
CREATE UNIQUE INDEX user_auths_email_unique_idx ON user_auths(email);

-- user_products stores users' products.
CREATE TABLE user_products (
  -- user_product_id represents a product id. This is UUID v4.
  user_product_id varchar(36) NOT NULL PRIMARY KEY,
  -- user_id represents an id from the users table.
  user_id varchar(36) NOT NULL,
  -- title represents a title.
  title varchar(255) NOT NULL,
  -- description represents a description.
  description varchar(2000) NOT NULL,
  -- create_at represents a creation timestamp.
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  -- update_at represents a last update timestamp.
  updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE Now(),
  CONSTRAINT FOREIGN KEY (user_id) REFERENCES users(user_id)
) ENGINE=InnoDB;

-- user_products_title_idx represents an index for title search.
CREATE INDEX user_products_title_idx ON user_products(title);