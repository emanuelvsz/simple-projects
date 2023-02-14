package handlers

import (
	"main/utils/data"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListGenres(c echo.Context) error {
	listGenres := data.Genres

	return c.JSON(http.StatusOK, listGenres)
}
