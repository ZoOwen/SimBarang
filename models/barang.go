package models

import "gorm.io/gorm"

type Barang struct {
	gorm.Model
	Nama       string `json:"nama"`
	SupplierId int16  `json:"supplier_id"`
	Harga      int64  `json:"harga"`
	CategoryId int16  `json:"category_id"`
	Qty        int64  `json:"qty"`
	Status     string `json:"status"`
}
