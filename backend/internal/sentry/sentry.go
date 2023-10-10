package sentry

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
)

func Setup() error {
	if debug() {
		return nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:                env.Vars.SentryDSN,
		Environment:        string(env.Vars.Stage),
		EnableTracing:      true,
		TracesSampleRate:   1.0,
		ProfilesSampleRate: 1.0,
	})
	if err != nil {
		return errors.Wrap(err, "failed to initialize sentry")
	}

	return nil
}

func Flush() {
	if debug() {
		return
	}

	sentry.Flush(10 * time.Second)
}

func CaptureExceptionWithGin(ctx *gin.Context, err error) {
	if debug() {
		return
	}

	hub := sentrygin.GetHubFromContext(ctx)
	if hub == nil {
		log.Warn(ctx, "failed to get sentry hub from context")
		return
	}
	_ = hub.CaptureException(err)
}

func debug() bool {
	return env.Vars.Stage == "local"
}
