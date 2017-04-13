package lib

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestSplitterBasic(t *testing.T) {
	var longTweet = `This is a long tweet that should be split into two separate tweets using the splitter function of this project. This should validate the basic splitting check and perform a simple functional test.`
	var scanner = bufio.NewScanner(strings.NewReader(longTweet))
	scanner.Split(ScanTweets)
	count := 0
	for scanner.Scan() {
		count++
		if txt := scanner.Text(); txt == "" {
			t.Errorf("invalid token text from split function")
		}
	}
	if count != 2 {
		t.Errorf("invalid tweet split, expected %d splits", 2)
	}
}

func TestSplitterFile(t *testing.T) {
	f, err := os.Open("story.txt")
	if err != nil {
		t.Error("error reading file")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(ScanTweets)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		t.Error("error during scanning tweets")
	}
	if count != 13 {
		t.Errorf("invalid tweet split, expected %d splits, got %d", 13, count)
	}
}
