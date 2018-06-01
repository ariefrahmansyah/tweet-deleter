package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// Tweet struct.
type Tweet struct {
	ID int64 `json:"id"`
}

func main() {
	accessToken := os.Getenv("accessToken")
	accessTokenSecret := os.Getenv("accessTokenSecret")
	consumerKey := os.Getenv("consumerKey")
	consumerSecret := os.Getenv("consumerSecret")

	api := anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret)
	log.Println("API Ready")

	files, err := ioutil.ReadDir("tweets")
	panicOnError(err)

	for _, f := range files {
		// Put json file from twitter archive in tweets folder
		file, err := os.Open("tweets/" + f.Name())
		panicOnError(err)

		log.Println(f.Name())

		l := 0
		lines := []byte{}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// SKip the first line from json file
			if l == 0 {
				l++
				continue
			}

			lines = append(lines, scanner.Bytes()...)
		}

		tweets := []Tweet{}
		err = json.Unmarshal(lines, &tweets)
		panicOnError(err)

		for _, tweet := range tweets {
			_, err := api.DeleteTweet(tweet.ID, true)
			if err != nil {
				log.Println("Failed to delete:", tweet.ID, err)
			}
		}
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
