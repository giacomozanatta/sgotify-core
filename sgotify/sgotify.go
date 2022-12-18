package sgotify

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"os"
)

const (
	AUTH_STATE = "sgotify"
)

func Client() (*spotify.Client, error) {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
		return nil, err
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	return spotify.New(httpClient), nil
}

func AuthURL(redirectURL string) string {
	fmt.Println(redirectURL)
	return spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURL),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate)).AuthURL(AUTH_STATE)
}
