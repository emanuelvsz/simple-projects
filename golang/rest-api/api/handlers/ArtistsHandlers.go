package handlers

import (
	"main/utils/data"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListArtists(c echo.Context) error {
	listArtist := data.Artists

	return c.JSON(http.StatusOK, listArtist)
}
