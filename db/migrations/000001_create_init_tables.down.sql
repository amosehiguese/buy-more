-- Delete Foreign Key Constraints First
ALTER TABLE product_categories DROP COLUMN IF EXISTS product_id;
ALTER TABLE product_categories DROP COLUMN IF EXISTS category_id;
ALTER TABLE images DROP COLUMN IF EXISTS product_id;

-- Then Delete Tables
DROP TABLE IF EXISTS product_categories;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS images;
