package domain

type Barang struct {
	Nama        string `json:"nama"`
	SupplierId  int16  `json:"supplier_id"`
	HargaGrosir int64  `json:"harga_grosir"`
	HargaSatuan int64  `json:"harga_satuan"`
	CategoryId  int16  `json:"category_id"`
	Qty         int64  `json:"qty"`
	Status      string `json:"status"`
}
