package handlers

// import (
// 	"main/api/utils"
// 	"main/models"
// 	"main/utils/data"
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

// func ListAlbums(context echo.Context) error {
// 	albumList := data.Albumns
// 	songList := data.Songs
// 	artistID, err := strconv.Atoi(context.Param(utils.ArtistID))
// 	if err != nil {
// 		artistID = 1
// 	}
// 	albumListNew := []models.Album{}

// 	for i := 0; i < len(albumList); i++ {
// 		for i := 0; i < len(songList); i++ {
// 			if songList[i].ArtistID == artistID {
// 				albumListNew[i].Songs = append(albumListNew[i].Songs, songList[i])
// 			}
// 		}
// 	}

// 	return context.JSON(http.StatusOK, albumListNew)
// }
