package handlers

import (
	"database/sql"
	"github.com/rtmelsov/simple-db/internal/database"
	"github.com/rtmelsov/simple-db/internal/models"
)

func AddAlbum(title, artist string, price float64) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", title, artist, price)
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
