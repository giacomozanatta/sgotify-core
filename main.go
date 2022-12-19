package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	"net/http"
	"sgotify-core/controllers"
	"sgotify-core/sgotify"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env not found")
	}
	var queue sgotify.Songs
	engine := gin.Default()
	engine.LoadHTMLGlob("client/templates/*")
	engine.Static("/css", "client/css")
	engine.StaticFile("/js/main.js", "client/scripts/main.js")
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/queue")
	})

	engine.GET("/queue", func(c *gin.Context) {
		c.HTML(http.StatusOK, "queue.html", gin.H{
			"SongsOnSpotify": queue.SongsOnSpotify(),
			"SongsOnQueue":   queue.SongsNotOnSpotify(),
		})
	})

	engine.GET("/search", func(c *gin.Context) {
		spotifyClient, err := sgotify.Client()
		if err != nil {
			fmt.Println("error")
		}
		searchPhrase := c.Query("q")
		searchResults, err := spotifyClient.Search(context.Background(), searchPhrase, spotify.SearchTypeTrack, spotify.Limit(20))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "NOT FOUND",
			})
		}
		var songs sgotify.Songs
		if searchResults.Tracks != nil {
			for _, item := range searchResults.Tracks.Tracks {
				artists := item.Artists[0].Name
				if len(item.Artists) > 1 {
					artists += ", " + item.Artists[1].Name
				}
				songs = append(songs, sgotify.Song{
					Title:     item.Name,
					Author:    artists,
					Id:        string(item.ID),
					InQueue:   queue.Contains(string(item.ID)),
					OnSpotify: queue.OnSpotify(string(item.ID)),
				})
			}
		}
		c.HTML(http.StatusOK, "search.html", gin.H{
			"Songs": songs,
			"Query": searchPhrase,
		})
	})

	engine.POST("/queue/add", func(c *gin.Context) {
		var song sgotify.Song

		if err := c.BindJSON(&song); err != nil {
			fmt.Println("ERROR")
			c.JSON(http.StatusOK, gin.H{
				"message": "",
			})
		}
		song.InQueue = true
		queue = append(queue, song)
		c.Writer.WriteHeader(http.StatusNoContent)
	})
	//go spotify_test()
	engine.GET("/admin/queue/new", func(c *gin.Context) {
		notOnSpotify := queue.SongsNotOnSpotify()
		c.JSON(http.StatusOK, gin.H{
			"songs": notOnSpotify,
		})
	})
	engine.PATCH("/admin/queue/:song_id", func(c *gin.Context) {
		var song sgotify.Song
		if err := c.BindJSON(&song); err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{
				"message": "",
			})
		}
		Id := c.Param("song_id")

		err := queue.SetOnSpotify(Id, song.OnSpotify)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{})
		}
		c.JSON(http.StatusOK, gin.H{})
	})

	engine.GET("/admin", controllers.GETAdmin)
	engine.GET("/admin/sgotipy/start", controllers.StartSgotipy)
	engine.GET("/admin/sgotipy/stop", controllers.StopSgotipy)
	engine.GET("spotify_auth_callback", controllers.CompleteSpotifyAuth)
	engine.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
