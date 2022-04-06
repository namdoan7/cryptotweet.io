package cmd

import (
	"context"
	"log"
	"os"

	"github.com/levinhne/cryptotweet.io/cmd/app"
	"github.com/levinhne/cryptotweet.io/cmd/app/command"
	"github.com/levinhne/cryptotweet.io/internal/tag/domain/tag"
	"github.com/urfave/cli/v2"
)

func Execute(app app.Application) error {
	tweetCommand := NewTweetCommand()
	profileCommand := NewProfileCommand()
	rootCommand := &cli.App{
		Name: "cryptotweet",
		Commands: []*cli.Command{
			{
				Name:  "tweet",
				Usage: "",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "",
						Before: func(c *cli.Context) error {
							// days := []string{}
							// prompt := &survey.MultiSelect{
							// 	Message: "What days do you prefer:",
							// 	Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
							// }
							// survey.AskOne(prompt, &days)
							return nil
						},
						Action: func(c *cli.Context) error {
							tweet, err := tweetCommand.GetTweetById(c.Args().First())
							app.Commands.CreateTweet.Handle(context.Background(), command.CreateTweet{
								Tweet: tweet,
							})
							return err
						},
						After: func(c *cli.Context) error {
							// log.Println("After", c.Context.Value("profile_id"))
							return nil
						},
					},
				},
			},
			{
				Name:  "profile",
				Usage: "",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "",
						Action: func(c *cli.Context) error {
							profile, err := profileCommand.GetProfileById(c.Args().First())
							log.Println(profile)
							app.Commands.CreateProfile.Handle(context.Background(), command.CreateProfile{
								Profile: profile,
							})
							return err
						},
						After: func(c *cli.Context) error {
							// log.Println("After", c.Context.Value("profile_id"))
							return nil
						},
					},
				},
			},
			{
				Name:  "tag",
				Usage: "",
				Subcommands: []*cli.Command{
					{
						Name:  "test",
						Usage: "",
						Action: func(c *cli.Context) error {
							tag, err := app.Commands.FinOrCreateTag.Handle(context.Background(), command.FindOrCreateTag{
								Tag: tag.Tag{
									Name: c.Args().First(),
								},
							})
							log.Println(tag)
							return err
						},
						After: func(c *cli.Context) error {
							// log.Println("After", c.Context.Value("profile_id"))
							return nil
						},
					},
				},
			},
		},
	}

	return rootCommand.Run(os.Args)
}
