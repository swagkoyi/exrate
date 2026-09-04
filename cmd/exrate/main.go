package main

import (
	"context"
	"fmt"
	"log"
	"os"

	api "exrate/internal/api"
	_ "exrate/internal/render"
	_ "exrate/internal/watchlist"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("action started")
			input := cmd.Args().Get(0)
			target := cmd.Args().Get(1)

			if input == "" || target == "" {
				log.Fatal("wrong type")
			}

			rate, err := api.GetRate(input, target)
			if err != nil {
				log.Fatal("wrong target")
			}
			fmt.Printf("%.8f", rate)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
