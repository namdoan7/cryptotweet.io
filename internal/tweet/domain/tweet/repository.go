package tweet

type Repository interface {
	Find() ([]Tweet, error)
	GetTweet(tweetId string) (*Tweet, error)
	Create(tweet Tweet) error
	Update(tweet Tweet) error
}
