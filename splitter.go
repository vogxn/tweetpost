package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	TWEET_CHAR_LIMIT int = 128
)

// ScanTweets is a split function for a Scanner that returns a sequence of
// tweets, each tweet not more than 128 characters in length.
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
			// if appending word breaks tweet char limit, finish scan
			word := data[start:i]
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
		return len(tweet), tweet[:], nil
	}
	// Request more data.
	return 0, nil, nil
}

func main() {
	f, err := os.Open("story.txt")
	if err != nil {
		fmt.Println("failed to read file %v", err)
		return
	}

	var scanner = bufio.NewScanner(f)
	scanner.Split(ScanTweets)
	count := 0
	for scanner.Scan() {
		count++
		fmt.Printf("(%#v/n) ", count)
		fmt.Println(scanner.Text())
		fmt.Println()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}
