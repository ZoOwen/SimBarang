package models

import "gorm.io/gorm"

type TransaksiReq struct {
	Transaksi    TransaksiBarang         `json:"transaksi"`
	DetailBarang []TransaksiBarangDetail `json:"detail_barang"`
}

type TransaksiBarang struct {
	gorm.Model
	Nota             string `json:"nota"`
	Nama             string `json:"nama"`
	Alamat           string `json:"alamat"`
	NoTelp           string `json:"no_telp"`
	TotalHarga       int64  `json:"total_harga"`
	TotalBayar       int64  `json:"total_bayar"`
	Kembalian        int64  `json:"kembalian"`
	StatusPembayaran string `json:"status_pembayaran"`
	Hutang           int64  `json:"hutang"`
}

type TransaksiRes struct {
	Transaksi    TransaksiBarang            `json:"transaksi"`
	DetailBarang []TransaksiBarangDetailRes `json:"detail_barang"`
}
type ResponeseGet struct {
	Status  int          `json:"status"`
	Data    TransaksiRes `json:"data"`
	Message string       `json:"message"`
}
