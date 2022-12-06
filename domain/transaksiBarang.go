package domain

type TransaksiReq struct {
	Transaksi    TransaksiBarang         `json:"transaksi"`
	DetailBarang []TransaksiBarangDetail `json:"detail_barang"`
}
type TransaksiBarang struct {
	Nota             string                  `json:"nota"`
	Nama             string                  `json:"nama"`
	Alamat           string                  `json:"alamat"`
	NoTelp           string                  `json:"no_telp"`
	TotalHarga       int64                   `json:"total_harga"`
	TipePembayaran   int64                   `json:"tipe_pembayaran"`
	TotalBayar       int64                   `json:"total_bayar"`
	Kembalian        int64                   `json:"kembalian"`
	StatusPembayaran string                  `json:"status_pembayaran"`
	DetailBarang     []TransaksiBarangDetail `json:"detail_barang,omitempty"`
}
type TransaksiBarangDetail struct {
	TransaksiId int64 `json:"transaksi_id"`
	IdBarang    int64 `json:"id_barang"`
	Harga       int64 `json:"harga"`
}
