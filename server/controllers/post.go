package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	tpost "github.com/vogxn/tweetpost"
	tpostlib "github.com/vogxn/tweetpost/lib"
	"github.com/vogxn/tweetpost/server/auth"
	"github.com/vogxn/tweetpost/server/views"
)

// read a maximum of 1MB of tweet data
const MAX_BODY int64 = 1024 * 1024

type PostHandle httprouter.Handle

var twitterAuth = auth.NewTwitterAuth("tweetpost")

// Homepage Index
func (ph *PostHandle) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Render the homepage template
	w.Header().Set("Content-Type", "text/html")
	var homePage = views.HomePage
	homePage.Render(w, nil)
}

// Post the tweet. We will do the OAuth flows here if the user is not authorized
func (ph *PostHandle) Tweet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var url string
	var err error
	defer r.Body.Close()

	if r.Method == "GET" && r.FormValue("oauth_token") != "" {
		// 2nd Leg of OAuth verification
		var verificationCode = r.FormValue("oauth_verifier")
		log.Println(twitterAuth.RequestToken)
		twitterAuth.AccessToken, err = twitterAuth.AuthorizeToken(twitterAuth.RequestToken, verificationCode)
		if err != nil {
			log.Fatal("error authorizing client")
		}
		log.Printf("AccessToken: %+v\n", twitterAuth.AccessToken)

		// Create HttpClient
		client, _ := twitterAuth.MakeHttpClient()
		response, _ := client.Get("https://api.twitter.com/1.1/statuses/home_timeline.json?count=1")
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		response.Write(w)

		// Set Form Values from Flash Messages
	} else if r.Method == "POST" {
		// Start OAuth here
		twitterAuth.RequestToken, url, err = twitterAuth.GetRequestTokenAndUrl("http://localhost:8080/tweet")
		if err != nil {
			log.Fatal(err)
		}

		// Redirect to twitter auth
		log.Println("Redirected to: ", url)
		log.Printf("RequestToken: %+v\n", twitterAuth.RequestToken)
		w.Header().Add("Access-Control-Allow-Origin", "*")
		http.Redirect(w, r, url, http.StatusFound)
	} else if r.Method == "OPTIONS" {
		log.Println("encountered OPTIONS header")
		w.WriteHeader(http.StatusOK)
	} else {
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Unable to perform OAuth")
	}
}

// Split input into tweets
func (ph *PostHandle) Split(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var post tpost.Post

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, MAX_BODY))
	if len(body) == 0 {
		log.Fatal("No post data recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Panicln(err)
	}
	// why are we doing this?
	if err := r.Body.Close(); err != nil {
		log.Panicln(err)
	}

	if err := json.Unmarshal(body, &post); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity) // Unprocessable Entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatal(err)
		}
	}

	// Success
	w.WriteHeader(http.StatusCreated)

	// Build tweet sequence
	var tweets []tpost.Tweet
	var scanner = bufio.NewScanner(strings.NewReader(post.Text))
	scanner.Split(tpostlib.ScanTweets)
	count := 0
	for scanner.Scan() {
		count++
		log.Println("scanned: ", scanner.Text())
		if count > 999 {
			log.Fatal("tweetstorm is too large than reserved space for sequence chars")
			w.WriteHeader(http.StatusInternalServerError)
		}
		seqTweet := fmt.Sprintf("%d/ %s", count, scanner.Text())
		tweets = append(tweets, tpost.Tweet{seqTweet})
	}

	if err := json.NewEncoder(w).Encode(tweets); err != nil {
		log.Fatal(err)
	}
	log.Println("Tweets: ", len(tweets))
}
