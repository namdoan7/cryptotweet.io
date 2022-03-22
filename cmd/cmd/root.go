package cmd

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func Execute() error {
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
							days := []string{}
							prompt := &survey.MultiSelect{
								Message: "What days do you prefer:",
								Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
							}
							survey.AskOne(prompt, &days)
							return nil
						},
						Action: func(c *cli.Context) error {
							_, err := tweetCommand.GetTweetById(c.Args().First())
							if err != nil {
								return err
							}
							// c.Context = context.WithValue(c.Context, "profile_id", t.TwitterProfileId)
							return nil
						},
						After: func(c *cli.Context) error {
							// log.Println("After", c.Context.Value("profile_id"))
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "",
						Action: func(c *cli.Context) error {
							return nil
						},
						After: func(c *cli.Context) error {
							// log.Println("After", c.Context.Value("profile_id"))
							return nil
						},
					},
					{
						Name:  "publish",
						Usage: "",
						Action: func(c *cli.Context) error {
							return nil
						},
						After: func(c *cli.Context) error {
							// log.Println("After", c.Context.Value("profile_id"))
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "",
						Action: func(c *cli.Context) error {
							return nil
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
							profileCommand.GetProfileById(c.Args().First())
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
