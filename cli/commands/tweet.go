package commands

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	tweetpb "github.com/cryptotweet.io/internal/common/genproto/tweet"
	"github.com/k0kubun/pp"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// đưa chormedp ra ngoài
// có thể đưa chromedp timeout
// đưa webview crawl ra ngoài
// spinner.CharSets

type TweetCommand struct {
}

func NewTweetCommand() *TweetCommand {
	return &TweetCommand{}
}

func (c *TweetCommand) AddTweet(tweetIds []string) error {
	// _ = color.New(color.Bold, color.FgGreen).SprintFunc()
	// spinner := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	// spinner.Color("bgBlack", "bold", "fgRed")
	// spinner.Suffix = " Started Google chrome"
	// spinner.Start()
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-infobars", true),
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
	requestIdChan := make(chan network.RequestID, len(tweetIds))
	chromedp.ListenTarget(ctx, func(v interface{}) {
		switch ev := v.(type) {
		case *network.EventResponseReceived:
			if strings.HasPrefix(ev.Response.URL, "https://cdn.syndication.twimg.com/tweet") {
				requestIdChan <- ev.RequestID
			}
		}
	})

	if err := chromedp.Run(ctx,
		chromedp.Navigate(`https://strongpasswordsgenerator.net/twitter.html?type=tweet&status=`+strings.Join(tweetIds, ",")),
	); err != nil {
		log.Fatal(err)
	}

	var data [][]byte
	for i := 0; i < len(tweetIds); i++ {
		requestId := <-requestIdChan
		if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
			body, err := network.GetResponseBody(requestId).Do(ctx)
			data = append(data, body)
			return err
		})); err != nil {
			log.Fatal(err)
		}
	}
	// https://stackoverflow.com/questions/64044242/how-to-convert-a-string-to-google-protobuf-timestamp
	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	// ccc := tweetpb.NewTweetServiceClient(conn)
	// ctxccc, _ := context.WithTimeout(context.Background(), time.Second)
	// ccc.Create(ctxccc, &tweetpb.CreateTweetRequest{
	// 	TweetId:          "",
	// 	TwitterProfileId: "",
	// 	Text: &tweetpb.Text{
	// 		En: "",
	// 		Vi: "",
	// 	},
	// 	ConversationCount: 0,
	// 	FavoriteCount:     0,
	// 	Lang:              "",
	// })

	// langs := []string{"vi", "ru"}
	done := make(chan bool)
	for _, d := range data {
		text, _ := jsonparser.GetString(d, "text")
		tweetId, _ := jsonparser.GetString(d, "id_str")
		tweetProfileId, _ := jsonparser.GetString(d, "user", "id_str")
		conversationCount, _ := jsonparser.GetInt(d, "conversation_count")
		favoriteCount, _ := jsonparser.GetInt(d, "favorite_count")
		lang, _ := jsonparser.GetString(d, "lang")
		createdAt, _ := jsonparser.GetString(d, "created_at")
		TweetedAt, _ := time.Parse(time.RFC3339, createdAt)
		a := &tweetpb.CreateTweetRequest{
			TweetId:          tweetId,
			TwitterProfileId: tweetProfileId,
			Text: &tweetpb.Text{
				En: text,
				Vi: "",
			},
			ConversationCount: int32(conversationCount),
			FavoriteCount:     int32(favoriteCount),
			Lang:              lang,
			TweetedAt:         timestamppb.New(TweetedAt),
		}

		pp.Println(a)

		// for _, l := range langs {
		// 	var s string
		// 	text, err := jsonparser.GetString(d, "text")
		// 	ctx2, cc := chromedp.NewContext(ctx)
		// 	chromedp.Run(ctx2,
		// 		chromedp.Navigate(`https://strongpasswordsgenerator.net/translate.html?text=`+url.QueryEscape(text)+`#googtrans/en/`+l),
		// 		chromedp.WaitVisible(`#text > font`),
		// 		chromedp.Text(`#text > font`, &s),
		// 		chromedp.ActionFunc(func(ctx context.Context) error {
		// 			time.Sleep(2 * time.Second)
		// 			cc()
		// 			return err
		// 		}),
		// 	)
		// }
	}
	<-done

	// done2 := make(chan bool)
	// var s string
	// langs := []string{"vi", "ru"}
	// i := 1

	// <-done2
	// spinner.Stop()
	// hashtag
	// category
	// Coin, Ex, Waleet, cryptocurrency market data
	// days := []string{}
	// prompt := &survey.MultiSelect{
	// 	Message: "What days do you prefer:",
	// 	Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
	// }
	// survey.AskOne(prompt, &days)
	// call grpc
	return nil
}
