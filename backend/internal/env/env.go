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
	Stage Stage `env:"STAGE,required"`
	// Slack
	SlackOAuthToken string `env:"SLACK_OAUTH_TOKEN,required"`
	// DynamoDB Table
	DynamoDBTableLGTMs   string `env:"DYNAMODB_TABLE_LGTMS,required"`
	DynamoDBTableReports string `env:"DYNAMODB_TABLE_REPORTS,required"`
	// S3 Bucket
	S3BucketImages string `env:"S3_BUCKET_IMAGES,required"`
	// SQS Queue
	SQSQueueNotifications    string `env:"SQS_QUEUE_NOTIFICATIONS,required"`
	SQSQueueURLNotifications string `env:"SQS_QUEUE_URL_NOTIFICATIONS,required"`
}

func Load() error {
	if err := env.Parse(&Vars); err != nil {
		return errors.Wrap(err, "failed to load env vars")
	}
	return nil
}
