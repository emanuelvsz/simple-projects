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
		Age: 27,
		Songs: []string{
			"As it was",
			"She",
			"Woman",
			"Late night talking",
			"Sign of the times",
			"Music for a sushi restaurant",
		},
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
		Age: 31,
		Songs: []string{
			"Run",
			"Sanctuary",
			"TEST DRIVE",
			"Can't get over you",
			"Slow dancing in the dark",
			"Glimpse of us",
		},
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
		Age: 31,
		Songs: []string{
			"Remember When",
			"Drunk on halloween",
			"Wish me luck",
			"These days",
			"Pictures of girls",
			"1980s horror film",
		},
	},
}
