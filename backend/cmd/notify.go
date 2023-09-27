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
	flagNotifyDebug  bool // --debug
)

var notifyCmd = &cobra.Command{
	Use: "notify",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !flagNotifyLambda {
			if !flagNotifyDebug {
				return errors.New("required --lambda or --debug")
			}

			ctx := context.Background()
			log.Info(ctx, "Debug mode")

			svc, err := service.New(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to create service")
			}
			err = svc.Notify(ctx, &events.SQSEvent{
				Records: []events.SQSMessage{
					{Body: `{
            "type": "lgtm_created",
            "lgtm_created": {
              "lgtm": {"id": "00000000-0000-0000-0000-000000000000"},
              "source": "https://example.com",
              "client_ip": "127.0.0.1"
            }
          }`},
				},
			})
			if err != nil {
				return errors.Wrap(err, "failed to notify")
			}

			return nil
		}

		lambda.Start(func(ctx context.Context, event *events.SQSEvent) error {
			svc, err := service.New(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to create service")
			}
			if err := svc.Notify(ctx, event); err != nil {
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
	notifyCmd.Flags().BoolVar(&flagNotifyDebug, "debug", false, "Run as debug")
}
