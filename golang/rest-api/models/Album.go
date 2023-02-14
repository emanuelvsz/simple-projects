package models

type Album struct {
	ID           int
	ArtistID     int
	Songs        []Song
	Colaborators []Artist
}
