package models

type Song struct {
	ID          int     `json:"id"`
	ArtistID    int     `json:"artist_id"`
	Name        string  `json:"name"`
	ReleaseDate string  `json:"release_date"`
	Duration    float32 `json:"duration"`
	Genre       []Genre `json:"genres"`
}
