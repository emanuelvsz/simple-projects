package data

import "main/models"

var Songs = []models.Song{
	{
		ID:          100,
		ArtistID:    1,
		Name:        "Late Night Talking",
		ReleaseDate: "2022",
		Duration:    3.42,
		Genre: []models.Genre{
			{
				ID:   1,
				Name: "Pop",
			},
			{
				ID:   2,
				Name: "Pop internacional",
			},
		},
	},
	{
		ID:          101,
		ArtistID:    2,
		Name:        "Run",
		ReleaseDate: "2020",
		Duration:    4.40,
		Genre: []models.Genre{
			{
				ID:   1,
				Name: "Pop",
			},
			{
				ID:   2,
				Name: "Pop Rock",
			},
		},
	},
	{
		ID:          102,
		ArtistID:    1,
		Name:        "Sign of the Times",
		ReleaseDate: "2017",
		Duration:    3.42,
		Genre: []models.Genre{
			{
				ID:   1,
				Name: "Pop",
			},
			{
				ID:   2,
				Name: "Pop internacional",
			},
		},
	},
	{
		ID:          103,
		ArtistID:    2,
		Name:        "TEST DRIVE",
		ReleaseDate: "2018",
		Duration:    2.34,
		Genre: []models.Genre{
			{
				ID:   1,
				Name: "Pop",
			},
			{
				ID:   2,
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
				ID:   1,
				Name: "Pop",
			},
			{
				ID:   2,
				Name: "Pop Indie",
			},
		},
	},
	{
		ID:          105,
		ArtistID:    1,
		Name:        "She",
		ReleaseDate: "2019",
		Duration:    3.42,
		Genre: []models.Genre{
			{
				ID:   1,
				Name: "Pop",
			},
			{
				ID:   2,
				Name: "Pop Rock",
			},
		},
	},
	{
		ID:          106,
		ArtistID:    1,
		Name:        "Falling",
		ReleaseDate: "2019",
		Duration:    3.42,
		Genre: []models.Genre{
			{
				ID:   1,
				Name: "Pop",
			},
			{
				ID:   2,
				Name: "Pop Rock",
			},
		},
	},
	{
		ID:          107,
		ArtistID:    3,
		Name:        "Pictures of Girls",
		ReleaseDate: "2018",
		Duration:    3.24,
		Genre: []models.Genre{
			{
				ID:   1,
				Name: "Indie",
			},
			{
				ID:   2,
				Name: "Indie Rock",
			},
		},
	},
	{
		ID:          108,
		ArtistID:    3,
		Name:        "Scrawny",
		ReleaseDate: "2019",
		Duration:    2.46,
		Genre: []models.Genre{
			{
				ID:   1,
				Name: "Indie",
			},
			{
				ID:   2,
				Name: "Indie Pop",
			},
		},
	},
	{
		ID:          109,
		ArtistID:    3,
		Name:        "Are you bored yet?",
		ReleaseDate: "2019",
		Duration:    2.58,
		Genre: []models.Genre{
			{
				ID:   1,
				Name: "Indie",
			},
			{
				ID:   2,
				Name: "Indie Pop",
			},
		},
	},
}
