package handlers

import (
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
	artistList := data.Artists

	id := c.Param("id")

	artistID, err := strconv.Atoi(id)

	if err != nil {
		artistID = 1
	}

	for i := 0; i < len(data.Artists); i++ {
		if artistList[i].ID == artistID {
			return c.JSON(http.StatusOK, artistList[i])
		}
	}

	return c.JSON(http.StatusNotFound, nil)
}
