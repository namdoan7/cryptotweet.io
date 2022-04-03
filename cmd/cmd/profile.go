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
	"github.com/levinhne/cryptotweet.io/internal/profile/domain/profile"
)

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

func (c *ProfleCommand) GetProfileById(profileId string) (*profile.Profile, error) {
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
		body, err := network.GetResponseBody(requestId).Do(ctx)
		if err != nil {
			return err
		}
		bodyJson := string(body)
		bodyJson = strings.ReplaceAll(bodyJson, "profile_image_url_https", "profile_image_url")
		bodyJson = strings.ReplaceAll(bodyJson, "pinned_tweet_ids_str", "pinned_tweet_ids")
		data, _, _, err = jsonparser.Get([]byte(bodyJson), "data", "user", "result") // legacy
		return err
	}))
	if err != nil {
		return nil, err
	}

	profileTwitterId, err := jsonparser.GetString(data, "rest_id")
	if err != nil {
		return nil, err
	}
	data, _, _, err = jsonparser.Get(data, "legacy")
	if err != nil {
		return nil, err
	}
	var profile profile.Profile
	err = json.Unmarshal(data, &profile)
	if err != nil {
		return nil, err
	}
	profile.ProfileTwitterId = profileTwitterId
	createdAt, err := jsonparser.GetString(data, "created_at")
	if err != nil {
		return nil, err
	}
	const RFC2822 = "Mon Jan 02 15:04:05 -0700 2006"
	profile.CreatedAt, err = time.Parse(RFC2822, createdAt)
	if err != nil {
		return nil, err
	}
	profile.UpdatedAt = profile.CreatedAt
	return &profile, err
}
