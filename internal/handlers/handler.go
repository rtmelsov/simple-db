package handlers

import (
	"github.com/rmelsov/simple-db/internal/database
	"github.com/rmelsov/simple-db/internal/models
)

func addAlbum(title, artist string, price float64) error {
	db := database.getDB()
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", title, artist, price)
	if err != nil {
		return err
	}
	if _, err = result.LastInsertId(); err != nil {
		return err
	}
	return nil
}

func getAlbumById(id int64) (*models.Album, error) {
	db := database.getDB()
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return a, nil
	}
	return a, nil
}

func getAlbumsByArtis(name string) ([]models.Album, error) {
	db := database.getDB()
	var albums []models.Album
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var alb Album
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
