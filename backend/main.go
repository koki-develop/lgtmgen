//go:generate go run github.com/swaggo/swag/cmd/swag fmt
//go:generate go run github.com/swaggo/swag/cmd/swag init --output ./swag

package main

import (
	"fmt"
	"os"

	"github.com/koki-develop/lgtmgen/backend/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("error: %+v\n", err)
		os.Exit(1)
	}
}
