package cmd

import (
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "cli",
}

func Execute() error {
	if err := env.Load(); err != nil {
		return err
	}

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
