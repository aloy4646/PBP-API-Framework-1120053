package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func TidakDipake() {
	//Ini cuma buat mysql bisa diimport
	x := mysql.ErrOldPassword
	fmt.Print(x)
}
