package data

import "main/models"

var Artists = []models.Artist{
	{
		ID:   1,
		Name: "Harry Styles",
		Genre: []string{
			"Indie pop",
			"Pop",
			"Pop melancólico",
		},
		Age:   27,
		Songs: []models.Song{},
	},
	{
		ID:   2,
		Name: "Joji",
		Genre: []string{
			"Indie pop",
			"Pop",
			"Pop melancólico",
			"Lo-fi",
		},
		Age:   31,
		Songs: []models.Song{},
	},
	{
		ID:   3,
		Name: "Wallows",
		Genre: []string{
			"Indie pop",
			"Indie rock",
			"Rock melancólico",
			"Indie",
		},
		Age:   31,
		Songs: []models.Song{},
	},
}
