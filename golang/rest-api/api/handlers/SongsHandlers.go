package handlers

import (
	"main/utils/data"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ListSongs(c echo.Context) error {
	listSongs := data.Songs

	return c.JSON(http.StatusOK, listSongs)
}

func ListSongByID(context echo.Context) error {
	songList := data.Songs

	id := context.Param("id")

	songID, err := strconv.Atoi(id)
	if err != nil {
		songID = 1
	}

	for i := 0; i < len(songList); i++ {
		if songList[i].ID == songID {
			return context.JSON(http.StatusOK, songList[i])
		}
	}

	return context.JSON(http.StatusNotAcceptable, nil)
}
