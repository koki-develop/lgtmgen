package cmd

import (
	"github.com/cockroachdb/errors"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "cli",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() error {
	if err := env.Load(); err != nil {
		return errors.Wrap(err, "failed to load env")
	}

	if err := rootCmd.Execute(); err != nil {
		return errors.Wrap(err, "failed to execute root command")
	}

	return nil
}
