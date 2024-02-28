package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main(){
	// capture conection properties
	cfg := mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "localhost",
		DBName: "recordings",
	}
	// get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil{
		fmt.Println("could not open database")
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Could not ping database")
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}