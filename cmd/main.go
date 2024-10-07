package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/phildehovre/go-complete-backend/cmd/api"
	"github.com/phildehovre/go-complete-backend/config"
	"github.com/phildehovre/go-complete-backend/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	initStorage(db)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)
	server.Run()

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database sucessfully connected")
}
