package tweetbook

import "time"

type Post struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

// Basic tweet struct that captures time of tweet and the
// author of the tweet
type Tweet struct {
	Text      string    `json:"text"`
	Author    string    `json:"author"`
	Timestamp time.Time `json:"timestamp"`
}
