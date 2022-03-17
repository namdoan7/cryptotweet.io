package main

import (
	"log"
	"os"
	"strings"

	"github.com/cryptotweet.io/cli/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	tweetCommand := commands.NewTweetCommand()
	// profileCommand := commands.NewProfileCommand()
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "profile",
				Usage: "add twitter profile",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new twitter profile",
						Action: func(c *cli.Context) error {
							return nil
							// return profileCommand.AddProfile(c.Args().First())
						},
					},
				},
			},
			{
				Name:  "tweet",
				Usage: "add twitter tweet",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new twitter tweet",
						Action: func(c *cli.Context) error {
							return tweetCommand.AddTweet(strings.Split(c.Args().First(), ","))
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
