package controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
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
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(
		os.Getenv("SPOTIFY_ID")+":"+os.Getenv("SPOTIFY_SECRET"))))
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}
