package data

import "main/models"

var Songs = []models.Song{
	{
		ID:          100,
		ArtistID:    1,
		Name:        "Late Night Talking",
		ReleaseDate: "10/12/2022",
		Duration:    3.42,
		Genre: []models.Genre{
			{
				ID: 1,
				Name: "Pop",
			},
			{
				ID: 2,
				Name: "Pop internacional",
			},
		},
	},
	{
		ID:          101,
		ArtistID:    2,
		Name:        "Run",
		ReleaseDate: "10/12/2020",
		Duration:    4.40,
		Genre: []models.Genre{
			{
				ID: 1,
				Name: "Pop",
			},
			{
				ID: 2,
				Name: "Pop Rock",
			},
		},
	},
	{
		ID:          102,
		ArtistID:    1,
		Name:        "Sign of the Times",
		ReleaseDate: "10/12/2022",
		Duration:    3.42,
		Genre: []models.Genre{
			{
				ID: 1,
				Name: "Pop",
			},
			{
				ID: 2,
				Name: "Pop internacional",
			},
		},
	},
	{
		ID:          103,
		ArtistID:    2,
		Name:        "TEST DRIVE",
		ReleaseDate: "10/12/2018",
		Duration:    2.34,
		Genre: []models.Genre{
			{
				ID: 1,
				Name: "Pop",
			},
			{
				ID: 2,
				Name: "Pop melancólico",
			},
		},
	},
	{
		ID:          104,
		ArtistID:    2,
		Name:        "Sanctuary",
		ReleaseDate: "10/12/2018",
		Duration:    3.34,
		Genre: []models.Genre{
			{
				ID: 1,
				Name: "Pop",
			},
			{
				ID: 2,
				Name: "Pop Indie",
			},
		},
	},
	{
		ID:          105,
		ArtistID:    1,
		Name:        "She",
		ReleaseDate: "10/12/2021",
		Duration:    3.42,
		Genre: []models.Genre{
			{
				ID: 1,
				Name: "Pop",
			},
			{
				ID: 2,
				Name: "Pop Rock",
			},
		},
	},
}
