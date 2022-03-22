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

type Url struct {
	DispayUrl   string  `bson:"display_url,omitempty" json:"display_url,omitempty"`
	ExpandedUrl string  `bson:"expanded_url,omitempty" json:"expanded_url,omitempty"`
	Indices     []uint8 `bson:"indices,omitempty" json:"indices,omitempty"`
	Url         string  `bson:"url,omitempty" json:"url,omitempty"`
}

type Description struct {
	Urls []Url `bson:"urls,omitempty" json:"urls,omitempty"`
}

type Entities struct {
	Description Description `bson:"description,omitempty" json:"description,omitempty"`
}

type Profile struct {
	Id               string    `bson:"_id,omitempty" json:"_id,omitempty"`
	ProfileTwitterId string    `bson:"profile_twitter_id,omitempty" json:"profile_twitter_id,omitempty"`
	Name             string    `bson:"name,omitempty" json:"name,omitempty"`
	ScreenName       string    `bson:"screen_name,omitempty" json:"screen_name,omitempty"`
	FavouritesCount  int32     `bson:"favourites_count,omitempty" json:"favourites_count,omitempty"`
	FollowersCount   int32     `bson:"followers_count,omitempty" json:"followers_count,omitempty"`
	Verified         bool      `bson:"verified,omitempty" json:"verified,omitempty"`
	Description      string    `bson:"description,omitempty" json:"description,omitempty"`
	ProfileImageUrl  string    `bson:"profile_image_url_https,omitempty" json:"profile_image_url_https,omitempty"`
	Entities         Entities  `bson:"entities,omitempty" json:"entities,omitempty"`
	PinnedTweetIds   []string  `bson:"pinned_tweet_ids,omitempty" json:"pinned_tweet_ids,omitempty"`
	CreatedAt        time.Time `bson:"created_at,omitempty" json:"-"`
	UpdatedAt        time.Time `bson:"updated_at,omitempty" json:"-"`
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
	createdAt, err := jsonparser.GetString(data, "created_at")
	if err != nil {
		return nil, err
	}
	const RFC2822 = "Mon Jan 02 15:04:05 -0700 2006"
	profile.CreatedAt, err = time.Parse(RFC2822, createdAt)
	if err != nil {
		return nil, err
	}
	return &profile, err
}
