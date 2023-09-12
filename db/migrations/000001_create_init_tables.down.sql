-- Delete Foreign Key Constraints First
ALTER TABLE images DROP COLUMN IF EXISTS product_id;
ALTER TABLE products DROP COLUMN IF EXISTS category_id;

-- Then Delete Tables
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS images;
