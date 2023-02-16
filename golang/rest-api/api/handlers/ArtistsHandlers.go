package handlers

import (
	"fmt"
	"main/utils/data"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ListArtists(c echo.Context) error {
	listArtist := data.Artists

	return c.JSON(http.StatusOK, listArtist)
}

func ListArtistByID(c echo.Context) error {
	id := c.QueryParam("artist/:artistID")

	artistID, err := strconv.Atoi(id)
	if err != nil {
		return http.ErrAbortHandler
	}

	for i := 0; i < len(data.Artists); i++ {
		if data.Artists[i].ID == artistID {
			artist := data.Artists[i]
			fmt.Println("Eoias")
			return c.JSON(http.StatusOK, artist)
		}
	}

	return c.JSON(http.StatusNotFound, nil)
}
