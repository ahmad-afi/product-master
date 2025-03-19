CREATE TABLE categories (
    id CHAR(26) PRIMARY KEY, -- ULID sebagai ID kategori
    name VARCHAR(100) NOT NULL, -- Nama kategori unik
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP not null,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP not null
);
CREATE INDEX idx_categories_deleted_at ON categories (deleted_at);
CREATE INDEX idx_categories_name ON categories (name);


INSERT INTO categories (id, name) VALUES
('01JPPW50103XYJHKYSVKHH30B5', 'Sayuran'),
('01JPPW5BPT88GK5WRGK26R5H1M', 'Protein'),
('01JPPW5H13C2053KNE5H1QPV5G', 'Buah'),
('01JPPW5NKYDAMAXGNG0NXGFYBB', 'Snack');

CREATE TABLE products (
    id CHAR(26) PRIMARY KEY, -- ULID sebagai ID produk
    name VARCHAR(255) NOT NULL,
    category_id CHAR(26) REFERENCES categories(id) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP not null,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP not null
);
INSERT INTO products (id, name, category_id, price) VALUES
('01JPPW5Z002B72060R3M639NV7', 'Bayam', '01JPPW50103XYJHKYSVKHH30B5', 5000),
('01JPPW62VGHBWNBR0YDQD68XZ1', 'Dada Ayam', '01JPPW5BPT88GK5WRGK26R5H1M', 30000),
('01JPPW6856YAJTBF91FE847869', 'Apel Fuji', '01JPPW5H13C2053KNE5H1QPV5G', 20000),
('01JPPW6BCS0QXPP6KY84ZDZ3KD', 'Keripik Kentang', '01JPPW5NKYDAMAXGNG0NXGFYBB', 15000);

CREATE INDEX idx_products_name ON products (name);
CREATE INDEX idx_products_category_id ON products (category_id);
CREATE INDEX idx_products_price ON products (price);
CREATE INDEX idx_products_created_at ON products (created_at DESC);
CREATE INDEX idx_products_deleted_at ON products (deleted_at);
