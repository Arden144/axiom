package search

import (
	"github.com/arden144/axiom/config"
	"github.com/arden144/axiom/log"
	"go.uber.org/zap"
)

type authorization struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func init() {
	var result authorization

	_, err := client.R().
		SetFormData(map[string]string{"grant_type": "client_credentials"}).
		SetBasicAuth(config.Spotify.ClientID, config.Spotify.ClientSecret).
		SetResult(&result).
		Post(AUTH)

	if err != nil {
		log.L.Fatal("failed to get spotify token", zap.Error(err))
	}

	client.SetCommonBearerAuthToken(result.AccessToken)
}
