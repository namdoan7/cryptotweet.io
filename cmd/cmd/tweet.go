package cmd

import (
	"context"
	"encoding/json"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/levinhne/cryptotweet.io/internal/tweet/domain/tweet"
)

type TweetCommand struct {
	ctx context.Context
}

func NewTweetCommand() *TweetCommand {
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
	return &TweetCommand{
		ctx: ctx,
	}
}

func (c *TweetCommand) GetTweetById(tweetId string) (*tweet.Tweet, error) {
	ctx, cancel := context.WithTimeout(c.ctx, 30*time.Second)
	defer cancel()
	done := make(chan bool)

	var requestId network.RequestID

	chromedp.ListenTarget(ctx, func(v interface{}) {
		switch ev := v.(type) {
		case *network.EventResponseReceived:
			if strings.HasPrefix(ev.Response.URL, "https://cdn.syndication.twimg.com/tweet") {
				requestId = ev.RequestID
				done <- true
			}
		}
	})

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://strongpasswordsgenerator.net/twitter.html?type=tweet&status=`+tweetId),
	)
	if err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
	case <-done:
	}
	var data string
	err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		body, err := network.GetResponseBody(requestId).Do(ctx)
		data = string(body)
		data = strings.ReplaceAll(data, "expandedUrl", "expanded_url")
		return err
	}))
	var tweet tweet.Tweet
	dataToBytes := []byte(data)
	err = json.Unmarshal(dataToBytes, &tweet)
	if err != nil {
		return nil, err
	}
	createdAt, err := jsonparser.GetString(dataToBytes, "created_at")
	if err != nil {
		return nil, err
	}
	tweetedAt, err := time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return nil, err
	}
	tweet.TweetedAt = tweetedAt
	tweetProfileId, err := jsonparser.GetString(dataToBytes, "user", "id_str")
	if err != nil {
		return nil, err
	}
	tweet.TwitterProfileId = tweetProfileId
	tweet.TweetId = tweetId
	langs := []string{"vi", "ru"}
	for _, lang := range langs {
		var result string
		ctx, _ = chromedp.NewContext(ctx)
		err = chromedp.Run(ctx,
			chromedp.Navigate(`https://strongpasswordsgenerator.net/translate.html?text=`+url.QueryEscape(tweet.Text)+`#googtrans/en/`+lang),
			chromedp.WaitVisible(`#text > font`),
			chromedp.Text(`#text > font`, &result),
			chromedp.ActionFunc(func(ctx context.Context) error {
				return err
			}),
		)
		if err != nil {
			return nil, err
		}
		switch lang {
		case "vi":
			tweet.TranslateText.Vietnamese = result
		case "ru":
			tweet.TranslateText.Russian = result
		}
	}
	return &tweet, err
}
