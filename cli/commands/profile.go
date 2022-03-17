package commands

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

type ProfileCommand struct {
}

func NewProfileCommand() *ProfileCommand {
	return &ProfileCommand{}
}

func (c *ProfileCommand) AddProfile(profileId string) error {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-features=IsolateOrigins", true),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("disable-site-isolation-trials", true),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
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

	// all we need to do here is navigate to the download url
	if err := chromedp.Run(ctx,
		chromedp.Navigate(`https://twitter.com/`+profileId),
	); err != nil {
		log.Fatal(err)
	}

	// This will block until the chromedp listener closes the channel
	<-done

	if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		_, err := network.GetResponseBody(requestId).Do(ctx)
		// var data map
		// pp.Println(jsonparser.GetInt(body, "data", "user", "result", "legacy", "fast_followers_count"))
		return err
	})); err != nil {
		log.Fatal(err)
	}
	log.Printf("Download Complete")
	return nil
}
