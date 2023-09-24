package cmd

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
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
		r := api.NewEngine()

		if flagServeLambda {
			lambda.Start(serveLambdaHandler(r))
			return nil
		}

		if err := r.Run(":8080"); err != nil {
			return err
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
