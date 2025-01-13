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
	objs, err := handlers.GetAlbumsByArtis("John Coltrane")
	obj, err := handlers.GetAlbumById(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("albums: %v\r\n", objs)
	fmt.Printf("album: %v\r\n", obj)

	err = handlers.AddAlbum("love is hurt", "Andi Topalidis", 56.56)
	if err != nil {
		log.Fatal(err)
	}
	var new []models.Album
	new, err = handlers.GetAlbumsByArtis("Andi Topalidis")
	fmt.Printf("new obj: %v\r\n", new)

	http.ListenAndServe(":8080", nil)
}
