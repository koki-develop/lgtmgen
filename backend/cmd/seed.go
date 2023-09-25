package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var seedCmd = &cobra.Command{
	Use: "seed",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		if env.Vars.Stage != env.StageLocal {
			return errors.New("seed command is only available in local stage")
		}

		type dummy struct {
			Color string
			Size  [2]int
		}

		dummies := []dummy{
			{Color: "000000", Size: [2]int{300, 300}},
			{Color: "ff0000", Size: [2]int{600, 600}},
			{Color: "00ff00", Size: [2]int{900, 900}},
			{Color: "0000ff", Size: [2]int{300, 600}},
			{Color: "ffff00", Size: [2]int{300, 900}},
			{Color: "00ffff", Size: [2]int{600, 300}},
			{Color: "ff00ff", Size: [2]int{600, 900}},
			{Color: "0fff00", Size: [2]int{900, 300}},
			{Color: "00fff0", Size: [2]int{900, 600}},
		}

		client := &http.Client{}
		for _, d := range dummies {
			p, err := json.Marshal(
				map[string]interface{}{
					"url": fmt.Sprintf("https://placehold.jp/%s/ffffff/%dx%d.png", d.Color, d.Size[0], d.Size[1]),
				},
			)
			if err != nil {
				return err
			}
			r := bytes.NewBuffer(p)

			url := "http://localhost:8080/v1/lgtms"
			log.Info(ctx, "request", "url", url, "body", string(p))
			req, err := http.NewRequest(http.MethodPost, url, r)
			if err != nil {
				return err
			}
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusCreated {
				err := errors.New("failed to create image")
				log.Error(ctx, "failed to create image", err, "status", resp.StatusCode)
				return err
			}

			time.Sleep(1 * time.Second)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
