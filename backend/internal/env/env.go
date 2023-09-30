package env

import (
	"github.com/caarlos0/env/v9"
	"github.com/cockroachdb/errors"
)

var (
	Vars Env
)

type Stage string

const (
	StageLocal      Stage = "local"
	StageDev        Stage = "dev"
	StageProduction Stage = "prd"
)

type Env struct {
	Stage          Stage  `env:"STAGE,required"`
	FrontendOrigin string `env:"FRONTEND_ORIGIN,required"`
	ImagesBaseURL  string `env:"IMAGES_BASE_URL,required"`
	// Google
	GoogleAPIKey   string `env:"GOOGLE_API_KEY,required"`
	SearchEngineID string `env:"SEARCH_ENGINE_ID,required"`
	// Slack
	SlackOAuthToken string `env:"SLACK_OAUTH_TOKEN,required"`
	// SQS Queue
	SQSQueueURLNotifications string `env:"SQS_QUEUE_URL_NOTIFICATIONS,required"`
}

func Load() error {
	if err := env.Parse(&Vars); err != nil {
		return errors.Wrap(err, "failed to load env vars")
	}
	return nil
}
