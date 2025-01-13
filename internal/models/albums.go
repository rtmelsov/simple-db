package albums

import "context"

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float64
}

type AlbumRepo interface {
	GetAlbums(context.Context) ([]Album, error)
	AddAlbum(context.Context, Album) (int, error)
}

