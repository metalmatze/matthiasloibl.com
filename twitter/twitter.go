package main

import (
	"errors"
	"log"

	"github.com/ChimeraCoder/anaconda"
	cache "github.com/patrickmn/go-cache"
)

type Tweet struct {
	ID        int64  `json:"id"`
	Text      string `json:"text"`
	Favorites int    `json:"favorites"`
	Retweets  int    `json:"retweets"`
}

type Twitter struct {
	API   *anaconda.TwitterApi
	Cache *cache.Cache
}

func (t Twitter) Tweets() ([]Tweet, error) {
	tweets, found := t.Cache.Get("tweets")
	if !found {
		timeline, err := t.API.GetUserTimeline(nil)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Can't get the users timeline")
		}

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
