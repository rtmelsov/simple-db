package handlers

import (
	"database/sql"
	"fmt"
	"github.com/rtmelsov/simple-db/internal/database"
	"github.com/rtmelsov/simple-db/internal/models"
)

func AddAlbum(album models.Album) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return err
	}
	if _, err = result.LastInsertId(); err != nil {
		return err
	}
	return nil
}

func GetAlbumById(id int64) (*models.Album, error) {
	db, err := database.Connect()
	var a models.Album
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	err = row.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return &a, nil
	}
	return &a, nil
}

func GetAlbumsByArtis(name string) ([]models.Album, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	var albums []models.Album
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var alb models.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, err
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}

func AddAlbums(books []models.Album) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, book := range books {
		res, err := stmt.Exec(book.Title, book.Artist, book.Price)
		if err != nil {
			return err
		}
		id, _ := res.LastInsertId()
		fmt.Println("Created book with ID", id)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GetAlbums() ([]models.Album, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	var albums []models.Album
	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var alb models.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, err
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}
