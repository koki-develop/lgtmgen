package cmd

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cockroachdb/errors"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/service"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	flagNotifyLambda bool // --lambda
)

var notifyCmd = &cobra.Command{
	Use: "notify",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !flagNotifyLambda {
			log.Info(context.Background(), "Only lambda is supported")
			return nil
		}

		lambda.Start(func(ctx context.Context, event *events.SQSEvent) error {
			svc, err := service.New(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to create service")
			}
			if err := svc.NotifyLGTMCreated(ctx, event); err != nil {
				return errors.Wrap(err, "failed to notify lgtm created")
			}
			return nil
		})

		return nil
	},
}

func init() {
	rootCmd.AddCommand(notifyCmd)
	notifyCmd.Flags().BoolVar(&flagNotifyLambda, "lambda", false, "Run as lambda")
}
