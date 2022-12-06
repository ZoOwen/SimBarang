package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	NoTelp string `json:"no_telp"`
}
