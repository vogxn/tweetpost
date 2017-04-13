package main

/* This will serve the REST calls for tweetbook */

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	tb "github.com/vogxn/tweetbook"
	tblib "github.com/vogxn/tweetbook/lib"
)

// read a maximum of 1MB of tweet data
const MAX_BODY int64 = 1024 * 1024

/* HomePage */
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Use the html/template library here
	fmt.Fprint(w, "Welcome to Tweetbook!\n")
}

/* Splitter POST */
func Split(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var post tb.Post
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, MAX_BODY))
	if err != nil {
		log.Panicln(err)
	}
	// why are we doing this?
	if err := r.Body.Close(); err != nil {
		log.Panicln(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.Unmarshal(body, &post); err != nil {
		w.WriteHeader(422) // Unprocessable Entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatal(err)
		}
	}

	// Success
	w.WriteHeader(http.StatusCreated)

	// Write tweet sequences
	var tweets []tb.Tweet
	var scanner = bufio.NewScanner(strings.NewReader(post.Text))
	scanner.Split(tblib.ScanTweets)
	count := 0
	for scanner.Scan() {
		count++
		log.Println("scanned: ", scanner.Text())
		tweets = append(tweets, tb.Tweet{scanner.Text(), post.Author, time.Now()})
	}

	if err := json.NewEncoder(w).Encode(tweets); err != nil {
		log.Fatal(err)
	}
	log.Println("Tweets: ", len(tweets))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/split", Split)

	log.Fatal(http.ListenAndServe(":8080", router))
}
