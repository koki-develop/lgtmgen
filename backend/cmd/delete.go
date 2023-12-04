package cmd

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cockroachdb/errors"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
	"github.com/koki-develop/lgtmgen/backend/internal/service"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	flagDeleteLambda bool // --lambda
)

var deleteCmd = &cobra.Command{
	Use: "delete",
	RunE: func(cmd *cobra.Command, args []string) error {
		if flagDeleteLambda {
			type Event struct {
				ID string `json:"id"`
			}

			lambda.Start(func(ctx context.Context, event Event) error {
				r, err := repo.New(ctx)
				if err != nil {
					return errors.Wrap(err, "failed to create repo")
				}
				svc, err := service.New(ctx, r)
				if err != nil {
					return errors.Wrap(err, "failed to create service")
				}

				if err := svc.DeleteLGTM(ctx, event.ID); err != nil {
					return errors.Wrap(err, "failed to delete lgtm")
				}

				return nil
			})
		} else {
			ctx := context.Background()

			if len(args) != 1 {
				return errors.New("invalid args")
			}
			id := args[0]

			r, err := repo.New(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to create repo")
			}
			svc, err := service.New(ctx, r)
			if err != nil {
				return errors.Wrap(err, "failed to create service")
			}

			if err := svc.DeleteLGTM(ctx, id); err != nil {
				return errors.Wrap(err, "failed to delete lgtm")
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&flagDeleteLambda, "lambda", "l", false, "Run as lambda")
}
