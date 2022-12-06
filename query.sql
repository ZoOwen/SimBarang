-- query barang --
INSERT INTO barangs (id, created_at, updated_at, deleted_at, nama, supplier_id, harga_grosir, harga_satuan, category_id, qty, status) VALUES 
		(NULL, NULL, NULL, NULL, 'indomie goreng', '5', '100000', '120000', '2', '50', 'ada'),
		(NULL, NULL, NULL, NULL, 'kopikap', '3', '80000', '90000', '3', '20', 'ada'),
        (NULL, NULL, NULL, NULL, 'sabun', '2', '50000', '60000', '5', '30', 'ada'),
        (NULL, NULL, NULL, NULL, 'teh gelas', '3', '70000', '80000', '7', '70', 'ada'),
        (NULL, NULL, NULL, NULL, 'chitato', '6', '180000', '190000', '2', '10', 'ada'),
        (NULL, NULL, NULL, NULL, 'sprite', '7', '100000', '120000', '8', '50', 'ada')

--query supplier --
INSERT INTO suppliers (id, created_at, updated_at, deleted_at, nama, alamat, no_telp) VALUES 
(NULL, NULL, NULL, NULL, 'PT. Raif membangun bangsa', 'Jalan tubagus ismail dalam', '0894230123412'),
(NULL, NULL, NULL, NULL, 'PT. fahmi membangun bangsa', 'Jalan Dipatiukur', '08942301234'),
(NULL, NULL, NULL, NULL, 'PT. Karunia membangun bangsa', 'Jalan Sekeloa', '089245698123'),
(NULL, NULL, NULL, NULL, 'PT. Indonesia membangun bangsa', 'Jalan raya pasteur', '089893762901');

-- query category --
INSERT INTO categories (name) VALUES
('Makanan'),
('Minuman'),
('Kesehatan'),
('Kebersihan'),
('Kebutuhan Harian')