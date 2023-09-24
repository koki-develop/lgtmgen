package cmd

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/api"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	flagServeLambda bool // --lambda
)

var serveCmd = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		r, err := api.NewEngine(ctx)
		if err != nil {
			return errors.Wrap(err, "failed to create engine")
		}

		if flagServeLambda {
			lambda.Start(serveLambdaHandler(r))
			return nil
		}

		if err := r.Run(":8080"); err != nil {
			return errors.Wrap(err, "failed to run server")
		}

		return nil
	},
}

type lambdaHandler func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func serveLambdaHandler(e *gin.Engine) lambdaHandler {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		ginLambda := ginadapter.New(e)
		return ginLambda.ProxyWithContext(ctx, req)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVar(&flagServeLambda, "lambda", false, "Run as lambda")
}
