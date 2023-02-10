package handlers

import (
	"main/utils/data"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListSongs(c echo.Context) error {
	ListSongs := data.Songs

	return c.JSON(http.StatusOK, ListSongs)
}
