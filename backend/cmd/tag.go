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
	flagTagLambda bool // --lambda
)

var tagCmd = &cobra.Command{
	Use: "tag",
	RunE: func(cmd *cobra.Command, args []string) error {
		if flagTagLambda {
			lambda.Start(func(ctx context.Context) error {
				r, err := repo.New(ctx)
				if err != nil {
					return errors.Wrap(err, "failed to create repo")
				}
				svc, err := service.New(ctx, r)
				if err != nil {
					return errors.Wrap(err, "failed to create service")
				}

				if err := svc.TagLGTM(ctx); err != nil {
					return errors.Wrap(err, "failed to tag LGTM")
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

			if err := svc.TagLGTM(ctx); err != nil {
				return errors.Wrap(err, "failed to tag LGTM")
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
	tagCmd.Flags().BoolVarP(&flagTagLambda, "lambda", "l", false, "Tag Lambda functions")
}
