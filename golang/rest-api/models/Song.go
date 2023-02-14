package models

type Song struct {
	ID          int
	ArtistID    int
	Name        string
	ReleaseDate string
	Duration    float32
	Genre       []Genre
}
