package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID int64
	Title string
	Artist string
	Price float32
}

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
	albums, err := albumByArtist("john Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}

func albumByArtist(name string)([]Album, error){
	// an album slic to hold data from returned rows
	var albums []Album

	rows, err := db.Query("select * from album where artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// loop through rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}