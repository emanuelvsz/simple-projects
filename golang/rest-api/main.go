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
	e.GET("/artist/:id", handlers.ListArtistByID)
	e.GET("/song", handlers.ListSongs)
	e.GET("/song/:id", handlers.ListSongByID)
	e.GET("/genre", handlers.ListGenres)
	e.GET("/genre/:id", handlers.ListGenreByID)

	e.Logger.Fatal(e.Start(":8082"))
}
