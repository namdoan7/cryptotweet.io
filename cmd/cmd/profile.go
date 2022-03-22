package cmd

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

type Profile struct {
	Id               string                 `bson:"_id,omitempty"`
	ProfileTwitterId string                 `bson:"profile_twitter_id,omitempty"`
	Name             string                 `bson:"name,omitempty"`
	ScreenName       string                 `bson:"screen_name,omitempty" json:"screen_name,omitempty"`
	FavouritesCount  int32                  `bson:"favourites_count,omitempty"`
	FollowersCount   int32                  `bson:"followers_count,omitempty"`
	Verified         bool                   `bson:"verified,omitempty"`
	Description      string                 `bson:"description,omitempty"`
	ProfileImageUrl  string                 `bson:"profile_image_url,omitempty"`
	Entities         map[string]interface{} `bson:"entities,omitempty"`
	PinnedTweetIds   []string               `bson:"pinned_tweet_ids,omitempty"`
	CreatedAt        time.Time              `bson:"created_at,omitempty"`
	UpdatedAt        time.Time              `bson:"updated_at,omitempty"`
}

type ProfleCommand struct {
	ctx context.Context
}

func NewProfileCommand() *ProfleCommand {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-features=IsolateOrigins", true),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("disable-site-isolation-trials", true),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	// create chrome instance
	ctx, _ := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	return &ProfleCommand{
		ctx: ctx,
	}
}

func (c *ProfleCommand) GetProfileById(profileId string) (*Profile, error) {
	ctx, cancel := context.WithTimeout(c.ctx, 30*time.Second)
	defer cancel()
	done := make(chan bool)

	var requestId network.RequestID

	chromedp.ListenTarget(ctx, func(v interface{}) {
		switch ev := v.(type) {
		case *network.EventResponseReceived:
			if strings.HasPrefix(ev.Response.URL, "https://twitter.com/i/api/graphql") && strings.Index(ev.Response.URL, "UserByScreenName") > -1 {
				requestId = ev.RequestID
				done <- true
			}
		}
	})

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://twitter.com/`+profileId),
	)
	<-done
	if err != nil {
		return nil, err
	}
	var data []byte
	err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		data, err = network.GetResponseBody(requestId).Do(ctx)
		if err != nil {
			return err
		}
		data, _, _, err = jsonparser.Get(data, "data", "user", "result", "legacy")
		return err
	}))

	var profile Profile
	err = json.Unmarshal(data, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, err
}
