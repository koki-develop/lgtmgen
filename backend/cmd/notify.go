package cmd

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cockroachdb/errors"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
	"github.com/koki-develop/lgtmgen/backend/internal/service"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	flagNotifyLambda bool // --lambda
	flagNotifyDebug  bool // --debug
	flagNotifyLGTM   bool // --lgtm
	flagNotifyReport bool // --report
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

			r, err := repo.New(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to create repo")
			}

			svc, err := service.New(ctx, r)
			if err != nil {
				return errors.Wrap(err, "failed to create service")
			}

			var body map[string]interface{}
			switch {
			case flagNotifyLGTM:
				body = map[string]interface{}{
					"type": "lgtm_created",
					"lgtm_created": map[string]interface{}{
						"lgtm": map[string]interface{}{
							"id": "00000000-0000-0000-0000-000000000000",
						},
						"source":    "https://example.com",
						"client_ip": "127.0.0.1",
					},
				}
			case flagNotifyReport:
				body = map[string]interface{}{
					"type": "report_created",
					"report_created": map[string]interface{}{
						"report": map[string]interface{}{
							"id":   "00000000-0000-0000-0000-000000000000",
							"type": "other",
							"text": "Hello, World!",
						},
					},
				}
			default:
				return errors.New("required --lgtm or --report")
			}

			p, err := json.Marshal(body)
			if err != nil {
				return errors.Wrap(err, "failed to marshal body")
			}
			err = svc.Notify(ctx, &events.SQSEvent{
				Records: []events.SQSMessage{{Body: string(p)}},
			})
			if err != nil {
				return errors.Wrap(err, "failed to notify")
			}

			return nil
		}

		lambda.Start(func(ctx context.Context, event *events.SQSEvent) error {
			log.Info(ctx, "received event", "event", event)

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
	notifyCmd.Flags().BoolVar(&flagNotifyLGTM, "lgtm", false, "Type is lgtm")
	notifyCmd.Flags().BoolVar(&flagNotifyReport, "report", false, "Type is report")
	notifyCmd.MarkFlagsMutuallyExclusive("lgtm", "report")
}
