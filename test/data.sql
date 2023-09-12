-- Categories
INSERT INTO categories (id, name, slug, created_at, updated_at)
VALUES
  ('3b3f6916-4f4d-4b7d-8d2c-2126efed3f64', 'Grocery', 'grocery', NOW(), NOW()),
  ('82f9c5ea-0ef6-4f3b-b2c6-b541efb848ab', 'Electronics', 'electronics', NOW(), NOW()),
  ('bb37e0d3-5d4d-4157-8ef4-4da46b084521', 'Clothing', 'clothing', NOW(), NOW());

-- Products
INSERT INTO products (id, category_id, name, slug, description, price, available, created_at, updated_at)
VALUES
  ('d8702d05-6f21-4c2b-b3b2-82e46c71e52e', '3b3f6916-4f4d-4b7d-8d2c-2126efed3f64', 'Tomatoes', 'tomatoes', 'Fresh and juicy tomatoes', '2.50', TRUE, NOW(), NOW()),
  ('35e348da-4963-4cda-88ef-9ad9bda0e0f9', '82f9c5ea-0ef6-4f3b-b2c6-b541efb848ab', 'Smartphone', 'smartphone', 'High-end smartphone with great features', '799.99', TRUE, NOW(), NOW()),
  ('6b9f3a3b-1f2c-4e5b-b1f7-4e7f6e181c4c', 'bb37e0d3-5d4d-4157-8ef4-4da46b084521', 'T-Shirt', 't-shirt', 'Comfortable cotton t-shirt', '19.99', TRUE, NOW(), NOW());

-- Images
INSERT INTO images (id, product_id, path, created_at, updated_at)
VALUES
  ('f07ae9c9-aa0e-4f01-a7b8-d5a8a64a59c3', 'd8702d05-6f21-4c2b-b3b2-82e46c71e52e', 'static/images/products/tomatoes.jpg', NOW(), NOW()),
  ('a9904a4f-e788-4b6b-95a8-45f68fcf319e', '35e348da-4963-4cda-88ef-9ad9bda0e0f9', 'static/images/products/smart-phone.jpg', NOW(), NOW()),
  ('8d5f0f5a-78d5-45cb-9c03-b7a31a6f0c6f', '6b9f3a3b-1f2c-4e5b-b1f7-4e7f6e181c4c', 'static/images/products/cotton-tee.jpg', NOW(), NOW());

