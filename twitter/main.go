package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/alexflint/go-arg"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/joho/godotenv"
	cache "github.com/patrickmn/go-cache"
)

type Config struct {
	ConsumerKey    string `arg:"env:CONSUMER_KEY"`
	ConsumerSecret string `arg:"env:CONSUMER_SECRET"`
	AccessToken    string `arg:"env:ACCESS_TOKEN"`
	AccessSecret   string `arg:"env:ACCESS_SECRET"`
}

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = level.NewFilter(logger, level.AllowDebug())
	level.Info(logger).Log("msg", "starting")

	_ = godotenv.Load()

	var config Config
	arg.MustParse(&config)

	c := cache.New(time.Minute, 10*time.Minute)

	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.AccessToken, config.AccessSecret)

	twitter := Twitter{
		API:    api,
		Cache:  c,
		Logger: logger,
	}

	http.HandleFunc("/", Handler(logger, twitter))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		level.Error(logger).Log("msg", "failed to listen and serve on :8080", "err", err)
	}
}

type Twitterer interface {
	Tweets() ([]Tweet, error)
}

func Handler(logger log.Logger, twitter Twitterer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tweets, err := twitter.Tweets()
		if err != nil {
			msg := "failed to get tweets"
			level.Warn(logger).Log("msg", msg, "err", err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		data, err := json.Marshal(tweets)
		if err != nil {
			msg := "failed to marshall tweets"
			level.Warn(logger).Log("msg", msg, "err", err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}
