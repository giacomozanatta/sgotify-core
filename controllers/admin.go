package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sgotify-core/sgotify"
	"sgotify-core/sgotipy"
)

func GETAdmin(c *gin.Context) {
	sgotipyStatus, err := sgotipy.GetStatus()
	if err != nil {
		fmt.Println(err.Error())
		sgotipyStatus = &sgotipy.SgotipyStatusResponse{
			Status:       "FAIL",
			Device:       "",
			DeviceStatus: "",
			CurrentSong:  nil,
		}
	}
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"SgotipyStatus": sgotipyStatus,
	})
	return

}

func StartSgotipy(c *gin.Context) {
	c.Redirect(http.StatusFound, sgotify.AuthURL(os.Getenv("SPOTIFY_AUTH_REDIRECT_URL")))
	return
}

func StopSgotipy(c *gin.Context) {
	err := sgotipy.StopSgotipy()
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusFound, "/admin")
	return
}
