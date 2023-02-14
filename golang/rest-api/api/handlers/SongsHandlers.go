package handlers

import (
	"main/utils/data"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListSongs(c echo.Context) error {
	listSongs := data.Songs

	return c.JSON(http.StatusOK, listSongs)
}
