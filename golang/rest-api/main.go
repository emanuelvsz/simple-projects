package main

import (
	"main/api/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/artist", handlers.ListArtists)
	e.GET("/artist/:artistID", handlers.ListArtistByID)
	e.GET("/artist/:artistID/song", handlers.ListSongsByArtistID)
	e.GET("/song", handlers.ListSongs)
	e.GET("/song/:songID", handlers.ListSongByID)
	e.GET("/genre", handlers.ListGenres)
	e.GET("/genre/:genreID", handlers.ListGenreByID)
	// e.GET("/artist/:artistID/album", handlers.ListAlbums)

	e.Logger.Fatal(e.Start(":8082"))
}
