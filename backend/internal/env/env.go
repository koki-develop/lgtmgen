package env

import (
	"github.com/caarlos0/env/v9"
)

var (
	Vars Env
)

type Env struct {
	// DynamoDB Table
	DynamoDBTableLGTMs string `env:"DYNAMODB_TABLE_LGTMS,required"`
}

func Load() error {
	if err := env.Parse(&Vars); err != nil {
		return err
	}
	return nil
}
