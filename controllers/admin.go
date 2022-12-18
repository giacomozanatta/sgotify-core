package controllers

import (
	"Sgotify/sgotify"
	"Sgotify/sgotipy"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func GETAdmin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"SpotifyAuthURL": sgotify.AuthURL(os.Getenv("SPOTIFY_AUTH_REDIRECT_URL")),
		"SgotipyStatus":  sgotipy.GetStatus(),
	})
}
