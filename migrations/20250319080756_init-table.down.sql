-- Hapus index pada tabel products
DROP INDEX IF EXISTS idx_products_name;
DROP INDEX IF EXISTS idx_products_category_id;
DROP INDEX IF EXISTS idx_products_price;
DROP INDEX IF EXISTS idx_products_created_at;
DROP INDEX IF EXISTS idx_products_deleted_at;

-- Hapus tabel products
DROP TABLE IF EXISTS products;

-- Hapus index pada tabel categories
DROP INDEX IF EXISTS idx_categories_name;
DROP INDEX IF EXISTS idx_categories_deleted_at;

-- Hapus tabel categories
DROP TABLE IF EXISTS categories;
