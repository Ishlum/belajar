package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:ahA2-BhH-BC-5dg5acC1cFchFEB5-AfB@tcp(godockerDB)/railway?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Database connected")
	DB = db
}
