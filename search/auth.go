package search

import (
	"context"
	"time"

	"github.com/arden144/axiom/bot"
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
	authorize()
}

func authorize() {
	var result authorization

	ctx, cancel := context.WithTimeout(bot.Ctx, 10*time.Second)
	defer cancel()

	_, err := client.R().
		SetContext(ctx).
		SetFormData(map[string]string{"grant_type": "client_credentials"}).
		SetBasicAuth(config.Spotify.ClientID, config.Spotify.ClientSecret).
		SetResult(&result).
		Post(AUTH)

	if err != nil {
		log.L.Fatal("failed to get spotify token", zap.Error(err))
	}

	client.SetCommonBearerAuthToken(result.AccessToken)

	time.AfterFunc(time.Duration(result.ExpiresIn-60)*time.Second, authorize)
}
