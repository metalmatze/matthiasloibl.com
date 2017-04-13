package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
	"github.com/go-errors/errors"
)

func main() {
	anaconda.SetConsumerKey("kq5CreMTj8gESKxNltU4i7LKK")
	anaconda.SetConsumerSecret("4KozKpXq3n4v56SneuTlRFBgagxC0sP30WWgeKaaS2lZrK3b45")
	api := anaconda.NewTwitterApi(
		"15984658-Vike4UqyZCxtOBn0uB3yjgm05JqgW1kUcyzWC7USg",
		"mCyZjI7gZpLU7Nf5OJZwUwqmD1sq0xfYnhbfdYpsOi2js",
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		timeline, err := timelineBytes(api)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(timeline)
	})

	log.Fatal(http.ListenAndServe(":45849", nil))
}

func timelineBytes(api *anaconda.TwitterApi) ([]byte, error) {

	timeline, err := api.GetUserTimeline(nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Can't get the users timeline")
	}

	timelineJson, err := json.Marshal(timeline)
	if err != nil {
		return nil, errors.New("Can't get json of the users timeline")
	}

	return timelineJson, nil
}
