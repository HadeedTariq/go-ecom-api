package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/HadeedTariq/go-ecom-api/cmd/api"
	"github.com/HadeedTariq/go-ecom-api/configs"
	"github.com/HadeedTariq/go-ecom-api/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	database, err := db.NewMySqlStorage(cfg)

	if err != nil {
		log.Fatal(err)
	}

	initMySqlStorage(database)

	server := api.NewApiServer(":8080", nil)

	fmt.Printf("Server listening on : %s", server.Addr)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initMySqlStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected successfully")
}
