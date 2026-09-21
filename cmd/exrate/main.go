package main

import (
	"context"
	"fmt"
	"log"
	"os"

	api "exrate/internal/api"
	render "exrate/internal/render"
	list "exrate/internal/watchlist"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "exrate",
		Usage: "write a list of currencies (exrate up)",
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
				Name:  "help",
				Usage: "--- exrate help\n			HAAAAAAAAAAAAAAAAALP",
				Action: func(ctx context.Context, c *cli.Command) error {
					root := c.Root()
					for _, sub := range root.Commands {
						fmt.Printf("%-15s %s\n", sub.Name, sub.Usage)
					}
					return nil
				},
			},
			{
				Name:  "cur",
				Usage: "--- exrate cur [input] [target]\n			exchange rate of the first currency against the second",
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
				Name:  "up",
				Usage: "--- exrate up {input, input, ...}\n			list of currencies; the list is displayed using `exrate` command",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("action 'up' started")
					watchlist := cmd.Args().Slice()
					return list.Save(watchlist)
				},
			},
			{
				Name:  "rend-list",
				Usage: "--- exrate rend-list [input]\n			display of exchange rates relative to the entered currency",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("action 'list' started")
					input := cmd.Args().First()
					if input == "" {
						log.Fatal("wrong cur")
					}
					rendr, err := api.Fetch(input)
					if err != nil {
						log.Fatal("wrong type")
					}
					result := render.RenderList(rendr)
					fmt.Println(result)
					return nil
				},
			},
			{
				Name:  "rend-sum",
				Usage: "--- exrate rend-sum [input] [target] [sum]\n			calculates the amount of the second currency relative to the first",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("action 'sum' started")
					input := cmd.Args().Get(0)
					target := cmd.Args().Get(1)
					num := cmd.Args().Get(2)
					one, err := api.GetRate(input, target)
					if err != nil {
						return err
					}
					two, err := render.RenderSum(one, num)
					if err != nil {
						return err
					}
					fmt.Println(two)
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
