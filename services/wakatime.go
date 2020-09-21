package services

import (
	"context"
	"os"

	"golang.org/x/oauth2"
)

// wakatimeから今日のコーディング時間合計を取得

type CredentialInfo struct {
	ClientID string
	ClientSecret string
	AuthorizeUrl string
	AccessTokenUrl string
	BaseUrl string	
}


func authenticate(){}

func GetTodaysCodingTime() () {
	clientId := os.Getenv("CLIENT_ID")
	clinetSecret := os.Getenv("CLIENT_SECRET")
	name := "wakatime"
	authorizeUrl := "https://wakatime.com/oauth/authorize"
	accessTokenUrl :="https://wakatime.com/oauth/token"
    baseUrl :="https://wakatime.com/api/v1/"


	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID: clientId,
		ClientSecret: clinetSecret,
		Scopes: []string{"email,read_stats"},
		Endpoint: oauth2.Endpoint{
			AuthURL: authorizeUrl,
			TokenURL: accessTokenUrl,
		},
	}

}
