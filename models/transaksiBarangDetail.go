package models

import "gorm.io/gorm"

type TransaksiBarangDetail struct {
	gorm.Model
	TransaksiId int   `json:"transaksi_id"`
	IdBarang    int   `json:"id_barang"`
	Harga       int64 `json:"harga"`
	JumlahBeli  int64 `json:"jumlah_beli"`
}

type TransaksiBarangDetailRes struct {
	gorm.Model
	TransaksiId int    `json:"transaksi_id"`
	Nota        string `json:"nota"`
	NamaBarang  string `json:"nama_barang"`
	Harga       int64  `json:"harga"`
	JumlahBeli  int64  `json:"jumlah_beli"`
}
