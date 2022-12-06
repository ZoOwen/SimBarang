package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zoowen/simbarang/config"
	"github.com/zoowen/simbarang/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//inisialisasi env
	username := os.Getenv("DBUSERNAME")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")
	dbHost := os.Getenv("DBHOST")
	//init koneksi dan migrate
	config.GormConnect(username, password, dbHost, dbName)
	config.Migrate()
	//init router
	router := routes.InitRouter()
	router.Run(os.Getenv("PORT"))
}
