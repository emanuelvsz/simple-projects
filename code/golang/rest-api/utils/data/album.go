package data

import "main/models"

var Albumns = []models.Album{
	{
		ID:           201,
		Name:         "Harry's House",
		ArtistID:     1,
		Songs:        []models.Song{},
		Colaborators: []models.Artist{},
	},
	{
		ID:           202,
		Name:         "Smithereens",
		ArtistID:     2,
		Songs:        []models.Song{},
		Colaborators: []models.Artist{},
	},
}
