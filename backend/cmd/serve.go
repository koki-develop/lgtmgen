package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {
		r := gin.Default()

		r.GET("/h", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})

		if err := r.Run(":8080"); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
