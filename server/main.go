package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vogxn/tweetpost/server/controllers"
)

func main() {

	var post controllers.PostHandle
	router := httprouter.New()
	// main routes
	router.GET("/", post.Index)
	router.POST("/split", post.Split)

	// final tweet poster
	router.POST("/tweet", post.Tweet)
	router.GET("/tweet", post.Tweet)

	// serve static assets
	router.GET("/static/*filepath", controllers.Static)
	log.Fatal(http.ListenAndServe(":8080", router))
}
