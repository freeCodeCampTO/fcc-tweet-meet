package main

import (
	"math/rand"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")

	log = &logger{logrus.New()}
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("did you forget your keys? " + name)
	}
	return v
}

func tweetFeed() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	api.SetLogger(log)

	rand.Seed(time.Now().UnixNano())

	meet := "every Saturday"

	texts := []string{
		"@freeCodeCampTO is meeting for coffee and code " + meet + "! Show up for friendly programming help from other developers anytime.", 
		"Stuck on a bug? Get coding help at @freeCodeCampTO, " + meet + ". ",
		"Need some pair programming help? Join @freeCodeCampTO " + meet + " for our weekly meetup!",
		"Having trouble coding? Get one-on-one help from friendly developers like yourself at @freeCodeCampTO, " + meet + ". ",
		"@freeCodeCampTO meets " + meet + ". Get live coding help from real people and meet fellow developers in Toronto!",
		"Join us @freeCodeCampTO, " + meet + " for our weekly coding session! Sometimes there's even cookies.",
		"Meet some friendly faces and learn to code with @freeCodeCampTO, " + meet + ". See you there!",
	}
	limit := len(texts)
	pick := rand.Intn(limit)

	tweet := texts[pick] + " https://freecodecampto.github.io #coding #programming #yyz"

	_, err := api.PostTweet(tweet, url.Values{})
	if err != nil {
		log.Critical(err)
	}
}

func main() {

	lambda.Start(tweetFeed)

}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
