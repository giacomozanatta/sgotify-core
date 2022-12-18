package controllers

import (
	"Sgotify/sgotify"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
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
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	spotifyAuth := sgotify.SpotifyAuth{}
	json.Unmarshal(body, &spotifyAuth)
	var conn *grpc.ClientConn
	conn, err = grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	sgotService := sgotify.NewSgotifyClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = sgotService.SendSpotifyAuth(ctx, &spotifyAuth)
	if err != nil {
		fmt.Println(err.Error())
	}
}
