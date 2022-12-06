package controllers

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"

	"github.com/zoowen/simbarang/config"
	"github.com/zoowen/simbarang/models"
	"github.com/zoowen/simbarang/util"
)

func Ping(ctx *gin.Context) {
	fmt.Println("hello")
	ctx.JSON(http.StatusCreated,
		gin.H{"response": http.StatusOK, "data": "pong", "message": "success"})

}

func AddTrxBarang(context *gin.Context) {
	var barang models.Barang
	var Transaksi models.TransaksiReq
	if err := context.ShouldBindJSON(&Transaksi); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var totalHarga int64
	for index, detBarang := range Transaksi.DetailBarang {

		id := detBarang.IdBarang
		db := config.GormDB.Table("barangs").Where("id = ? AND deleted_at is null", id).Scan(&barang)
		if err := db.Error; err != nil {
			util.LogError(err)
			context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		Transaksi.DetailBarang[index].Harga = barang.Harga
		totalHarga += detBarang.JumlahBeli * barang.Harga
		fmt.Println("total harga bayar", totalHarga)
		//update qty
		fmt.Println("qty barang", barang.Qty)
		UpdatedQty := barang.Qty - detBarang.JumlahBeli
		fmt.Println("total qty setelah di beli", UpdatedQty)
		updateQty := config.GormDB.Model(&barang).Update("qty", UpdatedQty).Where("id = ?", id)
		if err := updateQty.Error; err != nil {
			util.LogError(err)
			context.JSON(http.StatusInternalServerError, gin.H{"Message": "Error"})

		}
	}
	NOTA := "TRX" + string(time.Now().Format("02012006"))
	kembalian := Transaksi.Transaksi.TotalBayar - totalHarga
	Transaksi.Transaksi.TotalHarga = totalHarga
	Transaksi.Transaksi.Kembalian = kembalian
	Transaksi.Transaksi.Nota = NOTA
	if Transaksi.Transaksi.TotalBayar >= totalHarga {
		Transaksi.Transaksi.StatusPembayaran = "Lunas"
	} else {
		totalHutang := math.Abs(float64(kembalian))
		Transaksi.Transaksi.StatusPembayaran = "belum Lunas"
		Transaksi.Transaksi.Kembalian = 0
		Transaksi.Transaksi.Hutang = int64(totalHutang)
	}
	//insertData to Transaksi

	record := config.GormDB.Create(&Transaksi.Transaksi)
	if record.Error != nil {
		util.LogError(record.Error)
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error})
	}

	LastInsertId := Transaksi.Transaksi.ID

	//insert detail_barang to detail_barangs
	for index, data := range Transaksi.DetailBarang {
		data.TransaksiId = int(LastInsertId)
		insertBarang := config.GormDB.Create(&data)
		if insertBarang.Error != nil {
			util.LogError(insertBarang.Error)
			context.JSON(http.StatusInternalServerError, gin.H{"error": insertBarang.Error})
		}

		Transaksi.DetailBarang[index] = data
	}
	context.JSON(http.StatusCreated, gin.H{"response": http.StatusCreated, "data": Transaksi, "message": "success"})
}

func ListTransactions(ctx *gin.Context) {
	rdb := config.NewRedisClient()
	key := "ListTransaction-1"
	ttl := time.Duration(48) * time.Hour

	var data []models.TransaksiBarang
	if err := config.GormDB.Find(&data).Error; err != nil {
		ctx.AbortWithStatus(404)
		util.LogError(err)
	}

	dataRedis, _ := json.Marshal(data)
	// store data using SET command
	setToRedis := rdb.Set(context.Background(), key, dataRedis, ttl)
	if err := setToRedis.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
	log.Println("set operation success")
	getDataFromRedis := rdb.Get(context.Background(), key)
	if err := getDataFromRedis.Err(); err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	res, err := getDataFromRedis.Result()
	if err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	JsonData := []byte(res)
	dataUnmarshal := json.Unmarshal(JsonData, &data)
	// if errv2 != nil {
	// 	// panic
	// 	return
	// }
	fmt.Println("hasil json", dataUnmarshal)
	// fmt.Println("raw json", JsonData)
	log.Println("get operation success. result:", res)
	ctx.JSON(http.StatusOK, gin.H{"response": http.StatusOK, "data": data, "Message": "success"})
}

func GetDetailTransaction(context *gin.Context) {
	var TransaksiBarang models.TransaksiRes
	id := context.Param("id")

	if err := config.GormDB.Where("id = ?", id).First(&TransaksiBarang.Transaksi).Error; err != nil {
		context.AbortWithStatus(404)
		fmt.Println(err)
	}

	// var DetailBarang []models.TransaksiBarangDetail
	db := config.GormDB.Raw(`SELECT 
	td.id, td.transaksi_id,t.nota, b.nama as nama_barang, td.harga, td.jumlah_beli 
	from transaksi_barang_details td 
	left join transaksi_barangs t on t.id = td.transaksi_id
	left join barangs b on td.id_barang = b.id
	where td.transaksi_id = ? and td.deleted_at is null`, id).Scan(&TransaksiBarang.DetailBarang)

	if err := db.Error; err != nil {
		util.LogError(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": util.LogError(err)})
		return
	}
	if TransaksiBarang.Transaksi.ID == 0 {

		context.JSON(http.StatusAccepted, gin.H{"Message": "Error data tidak di temukan"})
	} else {
		context.JSON(http.StatusOK, gin.H{"response": 200, "data": TransaksiBarang, "message": "success"})

	}
}

func UpdateTransaction(context *gin.Context) {
	NOTA := "TRX" + string(time.Now().Format("02012006"))
	var TransaksiBarang models.TransaksiBarang
	TransaksiId := context.Param("id")
	db := config.GormDB.Table("transaksi_barangs").Where("id = ? AND deleted_at is null", TransaksiId).Scan(&TransaksiBarang)
	if err := db.Error; err != nil {
		fmt.Println("isi var err", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var Updateinput models.TransaksiBarang
	// var data models.TransaksiBarang
	if err := context.ShouldBindJSON(&Updateinput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Updateinput.Nota = NOTA
	fmt.Println("NOTA", NOTA)
	updated := config.GormDB.Model(&TransaksiBarang).Where("id = ? AND deleted_at is null", TransaksiId).Updates(Updateinput)
	if err := updated.Error; err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Error"})

	}
	fmt.Println(TransaksiBarang, TransaksiId, db)
	context.JSON(http.StatusOK, gin.H{"response": 200, "data": TransaksiBarang, "message": "success"})

}

// func DeleteTrx(context *gin.Context) {
// 	var publisher domain.Publisher
// 	id := context.Param("id")
// 	db := config.GormDB.Table("publishers").Where("id = ?", id).Delete(&publisher)
// 	if err := db.Error; err != nil {
// 		fmt.Println("isi var err", err)
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	context.JSON(http.StatusAccepted, gin.H{"Message": "Data berhasil di hapus!"})
// }
