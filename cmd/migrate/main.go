package main

import (
	"log"
	"os"

	"github.com/HadeedTariq/go-ecom-api/configs"
	"github.com/HadeedTariq/go-ecom-api/db"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
)

func main() {
	cfg := mysqlDriver.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// ~ so over there as this run out  side of the application so that's why I am registering the database with in the golang
	database, err := db.NewMySqlStorage(cfg)

	if err != nil {
		log.Fatal(err)
	}

	// ~ and here most probably integrating the migration library
	driver, err := mysqlMigrate.WithInstance(database, &mysqlMigrate.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	cmd := os.Args[len(os.Args)-1]

	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
