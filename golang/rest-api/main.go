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
	e.GET("/songs", handlers.ListSongs)
	e.GET("/genres", handlers.ListGenres)

	e.Logger.Fatal(e.Start(":8081"))
}

func artistsRoutes() {

}
