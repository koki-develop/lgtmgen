package sentry

import (
	"github.com/cockroachdb/errors"
	"github.com/getsentry/sentry-go"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
)

func Setup() error {
	if disabled() {
		return nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              env.Vars.SentryDSN,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		return errors.Wrap(err, "failed to initialize sentry")
	}

	return nil
}

func disabled() bool {
	return env.Vars.Stage == "local"
}
