package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
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
)

func main() {
	flag.Parse()

	if *configFile == "" || len(flag.Args()) == 0 {
		log.Fatal("Usage: tweet -c <config> <string to tweet>")
	}

	bytes, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal("Reading %s: %s", *configFile, err)
	}

	var conf Config
	if err = json.Unmarshal(bytes, &conf); err != nil {
		log.Fatal("Unmarshaling config: %s", err)
	}

	tweet := strings.Join(flag.Args(), " ")

	anaconda.SetConsumerKey(conf.ConsumerKey)
	anaconda.SetConsumerSecret(conf.ConsumerSecret)

	api := anaconda.NewTwitterApi(conf.AccessKey, conf.AccessSecret)
	_, err = api.PostTweet(tweet, nil)
	if err != nil {
		log.Fatal("Posting tweet: %s", err)
	}
}
