package cmd

import (
	"context"
	"os"

	"github.com/AlecAivazis/survey/v2"
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
		Usage: `
			██████╗██████╗ ██╗   ██╗██████╗ ████████╗ ██████╗ ████████╗██╗    ██╗███████╗███████╗████████╗
			██╔════╝██╔══██╗╚██╗ ██╔╝██╔══██╗╚══██╔══╝██╔═══██╗╚══██╔══╝██║    ██║██╔════╝██╔════╝╚══██╔══╝
			██║     ██████╔╝ ╚████╔╝ ██████╔╝   ██║   ██║   ██║   ██║   ██║ █╗ ██║█████╗  █████╗     ██║   
			██║     ██╔══██╗  ╚██╔╝  ██╔═══╝    ██║   ██║   ██║   ██║   ██║███╗██║██╔══╝  ██╔══╝     ██║   
			╚██████╗██║  ██║   ██║   ██║        ██║   ╚██████╔╝   ██║   ╚███╔███╔╝███████╗███████╗   ██║   
			╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝        ╚═╝    ╚═════╝    ╚═╝    ╚══╝╚══╝ ╚══════╝╚══════╝   ╚═╝   																							
		`,
		Commands: []*cli.Command{
			{
				Name:  "tweet",
				Usage: "",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "",
						Action: func(c *cli.Context) error {
							tweet, err := tweetCommand.GetTweetById(c.Args().First())
							hashtags := make([]string, 0)
							for _, hashtag := range tweet.Entities.Hashtags {
								hashtags = append(hashtags, hashtag.Text)
							}
							hts := []string{}
							hashtags = append(hashtags, []string{"ADA", "SHIBA"}...)
							prompt := &survey.MultiSelect{
								Message: "Choose a hashtag:",
								Options: hashtags,
								Default: "red",
							}
							survey.AskOne(prompt, &hts, survey.WithValidator(survey.Required))
							// for _, h := range hts {
							// 	app.Commands.
							// }
							// app.Commands.CreateTweet.Handle(context.Background(), command.CreateTweet{
							// 	Tweet: tweet,
							// })
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
							app.Commands.CreateTag.Handle(context.Background(), command.CreateTag{
								Tag: tag.Tag{
									Name: c.Args().First(),
								},
							})
							// tag, err := app.Commands.FinOrCreateTag.Handle(context.Background(), command.FindOrCreateTag{
							// 	Tag: tag.Tag{
							// 		Name: c.Args().First(),
							// 	},
							// })
							// log.Println(tag)
							// return err
							return nil
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
