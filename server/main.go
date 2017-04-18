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
	router.GET("/post", post.Index)
	router.POST("/post", post.Split)

	// serve static assets
	router.GET("/static/*filepath", controllers.Static)
	log.Fatal(http.ListenAndServe(":8080", router))
}
