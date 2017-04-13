package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/alexflint/go-arg"
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
	_ = godotenv.Load()

	var config Config
	arg.MustParse(&config)

	c := cache.New(time.Minute, 10*time.Minute)

	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.AccessToken, config.AccessSecret)

	twitter := Twitter{API: api, Cache: c}

	http.HandleFunc("/", Handler(twitter))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Twitterer interface {
	Tweets() ([]Tweet, error)
}

func Handler(twitter Twitterer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tweets, err := twitter.Tweets()
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to get tweets", http.StatusInternalServerError)
			return
		}
		data, err := json.Marshal(tweets)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to marshall tweets", http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}
