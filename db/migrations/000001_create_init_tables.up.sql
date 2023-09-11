-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create categories table
CREATE TABLE IF NOT EXISTS categories (
  id              UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  name            VARCHAR (200)  NOT NULL,
  slug            VARCHAR (200) NULL,
  created_at      TIMESTAMP DEFAULT NOW (),
  updated_at      TIMESTAMP NULL
);

-- Create products table
CREATE TABLE IF NOT EXISTS products (
  id              UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  category_id     UUID NOT NULL REFERENCES categories (id)  ON DELETE CASCADE,
  name            VARCHAR (200) NOT NULL,
  slug            VARCHAR (200) NULL,
  description     TEXT,
  price           DECIMAL (20, 2) NOT NULL,
  available       BOOLEAN NOT NULL,
  created_at      TIMESTAMP DEFAULT NOW (),
  updated_at      TIMESTAMP NULL
);

-- Create images table
CREATE TABLE IF NOT EXISTS images (
  id            UUID  DEFAULT uuid_generate_v4 () PRIMARY KEY,
  product_id    UUID NOT NULL REFERENCES products (id) ON DELETE CASCADE,
  path          VARCHAR (255) NOT NULL UNIQUE,
  created_at    TIMESTAMP DEFAULT NOW (),
  updated_at    TIMESTAMP NULL
);
