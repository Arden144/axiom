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

type failiure struct {
	error struct {
		status  int
		message string
	}
}

func init() {
	authorize()
}

func authorize() {
	var result authorization
	var fail failiure

	ctx, cancel := context.WithTimeout(bot.Ctx, 10*time.Second)
	defer cancel()

	res, err := client.R().
		SetContext(ctx).
		SetFormData(map[string]string{"grant_type": "client_credentials"}).
		SetBasicAuth(config.Spotify.ClientID, config.Spotify.ClientSecret).
		SetSuccessResult(&result).
		SetErrorResult(&fail).
		Post(AUTH)

	if err != nil || res.IsErrorState() {
		log.L.Error(
			"failed to renew spotify token",
			zap.Int("status", fail.error.status),
			zap.String("message", fail.error.message),
			zap.Error(err),
		)
		log.L.Info("trying to renew in 30s")
		time.AfterFunc(30*time.Second, authorize)
		return
	}

	client.SetCommonBearerAuthToken(result.AccessToken)
	log.L.Info("spotify token renewed", zap.Int("expires_in", result.ExpiresIn))

	time.AfterFunc(time.Duration(result.ExpiresIn-60)*time.Second, authorize)
}
