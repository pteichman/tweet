package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/url"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

type Config struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessKey      string
	AccessSecret   string
}

var (
	configFile = flag.String("c", "", "config file (json struct)")
	placeID    = flag.String("place", "", "place ID")
	lat        = flag.String("lat", "", "latitude")
	lon        = flag.String("lon", "", "longitude")
)

func main() {
	flag.Parse()

	if *configFile == "" || len(flag.Args()) == 0 {
		flag.PrintDefaults()
		log.Fatal("Missing config or text")
	}

	bytes, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("Reading %s: %s", *configFile, err)
	}

	var conf Config
	if err = json.Unmarshal(bytes, &conf); err != nil {
		log.Fatal("Unmarshaling config: %s", err)
	}

	tweet := strings.Join(flag.Args(), " ")

	anaconda.SetConsumerKey(conf.ConsumerKey)
	anaconda.SetConsumerSecret(conf.ConsumerSecret)

	args := url.Values{}
	if *placeID != "" {
		args["display_coordinates"] = []string{"true"}
		args["geo_enabled"] = []string{"true"}
		args["place_id"] = []string{*placeID}
	} else if *lat != "" && *lon != "" {
		args["display_coordinates"] = []string{"true"}
		args["geo_enabled"] = []string{"true"}
		args["lat"] = []string{*lat}
		args["lon"] = []string{*lon}
	}

	api := anaconda.NewTwitterApi(conf.AccessKey, conf.AccessSecret)
	_, err = api.PostTweet(tweet, args)
	if err != nil {
		log.Fatal("Posting tweet: %s", err)
	}
}
