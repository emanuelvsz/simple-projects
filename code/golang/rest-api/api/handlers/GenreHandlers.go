package handlers

import (
	"main/api/utils"
	"main/utils/data"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ListGenres(c echo.Context) error {
	listGenres := data.Genres

	return c.JSON(http.StatusOK, listGenres)
}

func ListGenreByID(context echo.Context) error {
	genreList := data.Genres

	id := context.Param(utils.GenreID)

	genreID, err := strconv.Atoi(id)
	if err != nil {
		genreID = 1
	}

	for i := 0; i < len(genreList); i++ {
		if genreList[i].ID == genreID {
			return context.JSON(http.StatusOK, genreList[i])
		}
	}

	return context.JSON(http.StatusNotAcceptable, nil)
}
