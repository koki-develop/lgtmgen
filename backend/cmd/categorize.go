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
	flagCategorizeLambda bool // --lambda
)

var categorizeCmd = &cobra.Command{
	Use: "categorize",
	RunE: func(cmd *cobra.Command, args []string) error {
		if flagCategorizeLambda {
			lambda.Start(func(ctx context.Context) error {
				r, err := repo.New(ctx)
				if err != nil {
					return errors.Wrap(err, "failed to create repo")
				}
				svc, err := service.New(ctx, r)
				if err != nil {
					return errors.Wrap(err, "failed to create service")
				}

				if err := svc.CategorizeLGTM(ctx); err != nil {
					return errors.Wrap(err, "failed to categorize LGTM")
				}

				return nil
			})
		} else {
			ctx := context.Background()

			r, err := repo.New(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to create repo")
			}
			svc, err := service.New(ctx, r)
			if err != nil {
				return errors.Wrap(err, "failed to create service")
			}

			if err := svc.CategorizeLGTM(ctx); err != nil {
				return errors.Wrap(err, "failed to categorize LGTM")
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(categorizeCmd)
	categorizeCmd.Flags().BoolVarP(&flagCategorizeLambda, "lambda", "l", false, "Run as lambda")
}
