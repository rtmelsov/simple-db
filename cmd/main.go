package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rmelsov/simple-db/internal/models"
	"github.com/rmelsov/simple-db/internal/handlers"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal()
	}
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db is connected")
	defer db.Close()
	var alb Album
	objs, err := alb.getAlbumsByArtis("John Coltrane")
	obj, err := alb.getAlbumById(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("albums: %v\r\n", objs)
	fmt.Printf("album: %v\r\n", obj)

	err = alb.addAlbum("love is hurt", "Andi Topalidis", 56.56)
	if err != nil {
		log.Fatal(err)
	}
	var new []Album
	new, err = alb.getAlbumsByArtis("Andi Topalidis")
	fmt.Printf("new obj: %v\r\n", new)

	http.ListenAndServe(":8080", nil)
}
