package tweetpost

type Post struct {
	Text string `json:"text"`
}

// Basic tweet struct that captures time of tweet and the
// author of the tweet
type Tweet struct {
	Text string `json:"text"`
}
