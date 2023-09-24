package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("serve called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
