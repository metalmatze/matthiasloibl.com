package main

import (
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	cache "github.com/patrickmn/go-cache"
)

type Tweet struct {
	ID        int64  `json:"id"`
	Text      string `json:"text"`
	Favorites int    `json:"favorites"`
	Retweets  int    `json:"retweets"`
}

type Twitter struct {
	API    *anaconda.TwitterApi
	Cache  *cache.Cache
	Logger log.Logger
}

func (t Twitter) Tweets() ([]Tweet, error) {
	tweets, found := t.Cache.Get("tweets")
	if !found {
		level.Debug(t.Logger).Log("msg", "tweets not found in in-memory cache")

		start := time.Now()
		timeline, err := t.API.GetUserTimeline(nil)
		if err != nil {
			return nil, err
		}
		level.Debug(t.Logger).Log("msg", "fetched tweets from twitter", "duration", time.Since(start))

		var tweets []Tweet
		for _, t := range timeline {
			tweets = append(tweets, Tweet{
				ID:        t.Id,
				Text:      t.Text,
				Favorites: t.FavoriteCount,
				Retweets:  t.RetweetCount,
			})
		}

		t.Cache.Set("tweets", tweets, cache.DefaultExpiration)
		return tweets, nil
	}

	return tweets.([]Tweet), nil
}
