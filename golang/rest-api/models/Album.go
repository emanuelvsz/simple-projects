package models

type Album struct {
	ID           int      `json:"id"`
	ArtistID     int      `json:"artist_id"`
	Songs        []Song   `json:"songs"`
	Colaborators []Artist `json:"colaborators"`
}
