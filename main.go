package main

import (
	"Sgotify/sgotify"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"log"
	"net/http"
)

const redirectURI = "http://localhost:8080/spotify_auth_callback"

var (
	_    = godotenv.Load()
	auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate,
		spotifyauth.ScopeUserReadPlaybackState,
		spotifyauth.ScopeStreaming))
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func spotify_test() {
	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)
	//fmt.Println(os.Getenv("SPOTIFY_ID"))
	//os.Setenv("SPOTIFY_ID", os.Getenv("SPOTIFY_ID"))
	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)
}
func main() {
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

	/*engine.GET("/spotify_auth_callback", func(c *gin.Context) {
		//completeAuth(c.Writer, c.Request)
	})*/
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
				fmt.Println("   ", item.Artists)
				fmt.Println(queue)
				fmt.Println(queue.Contains(string(item.ID)))
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
	engine.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

/*func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// use the token to get an authenticated client
	client := spotify.New(auth.Client(r.Context(), tok))
	//fmt.Fprintf(w, "Login Completed!")
	user, err := client.CurrentUser(r.Context())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(user)
	devices, _ := client.PlayerDevices(r.Context())
	fmt.Println(devices)
	ch <- client
}*/
