package env

import (
	"github.com/caarlos0/env/v9"
)

var (
	Vars Env
)

type Env struct {
	Stage string `env:"STAGE,required"`
	// DynamoDB Table
	DynamoDBTableLGTMs string `env:"DYNAMODB_TABLE_LGTMS,required"`
	// S3 Bucket
	S3BucketImages string `env:"S3_BUCKET_IMAGES,required"`
}

func Load() error {
	if err := env.Parse(&Vars); err != nil {
		return err
	}
	return nil
}
