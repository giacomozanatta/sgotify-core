package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"os"
	"sgotify-core/sgotipy"
	"strings"
)

func CompleteSpotifyAuth(c *gin.Context) {
	// use the token to get an authenticated client
	data := url.Values{
		"grant_type":   []string{"authorization_code"},
		"code":         []string{c.Query("code")},
		"redirect_uri": []string{os.Getenv("SPOTIFY_AUTH_REDIRECT_URL")},
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(
		os.Getenv("SPOTIFY_ID")+":"+os.Getenv("SPOTIFY_SECRET"))))
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	spotifyAuth := sgotipy.StartSgotipyRequest{}
	err = json.Unmarshal(body, &spotifyAuth)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = sgotipy.StartSgotipy(spotifyAuth)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusFound, "/admin")
	return
}
