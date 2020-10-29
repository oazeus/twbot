package twitterscraper

import (
	"context"
	"testing"
)

func TestGetSearchTweets(t *testing.T) {
	count := 0
	maxTweetsNbr := 50
	for tweet := range SearchTweets(context.Background(), "twitter scraper data -filter:retweets", maxTweetsNbr) {
		if tweet.Error != nil {
			t.Error(tweet.Error)
		} else {
			count++
			if tweet.HTML == "" {
				t.Error("Expected tweet HTML is not empty")
			}
			if tweet.ID == "" {
				t.Error("Expected tweet ID is not empty")
			}
			if tweet.PermanentURL == "" {
				t.Error("Expected tweet PermanentURL is not empty")
			}
			if tweet.Text == "" {
				t.Error("Expected tweet Text is not empty")
			}
		}
	}

	if count != maxTweetsNbr {
		t.Errorf("Expected tweets count=%v, got: %v", maxTweetsNbr, count)
	}
}
