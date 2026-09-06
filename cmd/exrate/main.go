package main

import (
	"context"
	"fmt"
	"log"
	"os"

	api "exrate/internal/api"
	_ "exrate/internal/render"
	list "exrate/internal/watchlist"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name: "exrate",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("action 'main' started")
			argsCount := cmd.Args().Len()
			if argsCount == 0 {
				listz, err := list.Load()
				if err != nil {
					return err
				} else {
					rates, err := api.GetRates(listz[0], listz[1:])
					if err != nil {
						return err
					}
					for _, currency := range listz[1:] {
						fmt.Printf("%s: %.5f\n", currency, rates[currency])
					}
				}
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "cur",
				Usage: "exrate cur [input] [target]",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("action 'cur' started")
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
			},
			{
				Name: "up",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("action 'up' started")
					watchlist := cmd.Args().Slice()
					return list.Save(watchlist)
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
