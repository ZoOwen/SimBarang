package config

import (
	"fmt"
	"log"

	"github.com/zoowen/simbarang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB *gorm.DB
var GormErr error

func GormConnect(username, password, dbHost, dbName string) {
	log.Println("Connected to Database!")
	strngConnection := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, dbHost, dbName)
	fmt.Println("ini koneksi", strngConnection)
	GormDB, GormErr = gorm.Open(mysql.Open(strngConnection), &gorm.Config{})
	if GormErr != nil {
		fmt.Println(GormErr.Error())
		panic("Cannot connect to DB")
	}

}

func Migrate() {
	GormDB.Migrator().DropTable(&models.Barang{},
		&models.Category{},
		&models.Supplier{})
	GormDB.AutoMigrate(
		&models.Barang{},
		&models.Category{},
		&models.Supplier{},
		&models.TransaksiBarang{},
		&models.TransaksiBarangDetail{},
		&models.User{},
	)
	GormDB.Exec(
		`INSERT INTO barangs (id, created_at, updated_at, deleted_at, nama, supplier_id, harga, category_id, qty, status) VALUES 
		(NULL, NULL, NULL, NULL, 'indomie goreng', '5', '100000', '2', '50', 'ada'),
		(NULL, NULL, NULL, NULL, 'kopikap', '3', '80000', '3', '20', 'ada'),
        (NULL, NULL, NULL, NULL, 'sabun', '2', '50000',  '5', '30', 'ada'),
        (NULL, NULL, NULL, NULL, 'teh gelas', '3', '80000', '7', '70', 'ada'),
        (NULL, NULL, NULL, NULL, 'chitato', '6',  '190000', '2', '10', 'ada'),
        (NULL, NULL, NULL, NULL, 'sprite', '7', '120000', '8', '50', 'ada')`,
	)
	GormDB.Exec(
		`INSERT INTO suppliers (id, created_at, updated_at, deleted_at, nama, alamat, no_telp) VALUES 
		(NULL, NULL, NULL, NULL, 'PT. Raif membangun bangsa', 'Jalan tubagus ismail dalam', '0894230123412'),
		(NULL, NULL, NULL, NULL, 'PT. fahmi membangun bangsa', 'Jalan Dipatiukur', '08942301234'),
		(NULL, NULL, NULL, NULL, 'PT. Karunia membangun bangsa', 'Jalan Sekeloa', '089245698123'),
		(NULL, NULL, NULL, NULL, 'PT. Indonesia membangun bangsa', 'Jalan raya pasteur', '089893762901');`,
	)
	GormDB.Exec(
		`INSERT INTO categories (nama) VALUES
		('Makanan'),
		('Minuman'),
		('Kesehatan'),
		('Kebersihan'),
		('Kebutuhan Harian')`,
	)
	log.Println("Database Migration Completed!")

}
