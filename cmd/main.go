package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/rtmelsov/simple-db/internal/database"
	"github.com/rtmelsov/simple-db/internal/handlers"
	"github.com/rtmelsov/simple-db/internal/models"
)

var db *sql.DB

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db is connected")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	albums := []models.Album{
		{Title: "love is hurt", Artist: "Andi Topalidis", Price: 56.56},
		{Title: "graduation", Artist: "Ye", Price: 33.56},
		{Title: "I'm stay", Artist: "Slava KPSS", Price: 74.56},
		{Title: "Hello", Artist: "Adel", Price: 96.56},
		{Title: "5 am in Piter", Artist: "Jubilee", Price: 199.56},
	}
	err = handlers.AddAlbum(albums[0])
	if err != nil {
		log.Fatal(err)
	}
	var new []models.Album
	new, err = handlers.GetAlbumsByArtis("Andi Topalidis")
	fmt.Printf("new obj: %v\r\n", new)

	err = handlers.AddAlbums(albums[1:])
	albs, err := handlers.GetAlbums()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("albums: %v", albs)

	http.ListenAndServe(":8080", nil)
}
