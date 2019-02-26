CREATE TABLE users (
  user_id integer NOT NULL AUTO_INCREMENT PRIMARY KEY
) ENGINE=InnoDB;

CREATE TABLE user_products (
  user_product_id integer NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id integer NOT NULL,
  title varchar(255) NOT NULL,
  description varchar(2000) NOT NULL,
  CONSTRAINT FOREIGN KEY (user_id) REFERENCES users(user_id)
) ENGINE=InnoDB;

CREATE INDEX user_products_title_idx ON user_products(title);