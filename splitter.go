package tweetbook

import (
	"unicode"
	"unicode/utf8"
)

const (
	TWEET_CHAR_LIMIT int = 140
)

// ScanTweets is a split function for a Scanner that returns a sequence of
// tweets, each tweet not more than TWEET_CHAR_LIMIT characters in length.
func ScanTweets(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	tweet := make([]byte, 0, TWEET_CHAR_LIMIT)
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}
	// Scan until tweet char limit
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if unicode.IsSpace(r) {
			// replace newlines with space mid-tweet
			switch data[i] {
			case '\n', '\r':
				data[i] = ' '
			}
			word := data[start:i]
			// if appending word breaks tweet char limit, finish scan
			if len(tweet)+len(word) >= TWEET_CHAR_LIMIT {
				return len(tweet) + width, tweet[:], nil
			} else {
				// append word to tweet
				tweet = append(tweet, word...)
				start = i
			}
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated chars. Return it.
	if atEOF && len(data) > start {
		return len(data), append(tweet, data[start:]...), nil
	}
	// Request more data.
	return 0, nil, nil
}
