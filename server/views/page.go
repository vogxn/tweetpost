package views

import (
	"html/template"
	"log"
	"net/http"

	"github.com/vogxn/tweetpost/lib"
)

type Page struct {
	Template *template.Template
	Layout   string
}

var HomePage *Page
var TweetPage *Page

// Render the layout as a html/template
func (page *Page) Render(w http.ResponseWriter, data interface{}) error {
	return page.Template.ExecuteTemplate(w, page.Layout, data)
}

func init() {
	homePageFiles := append(lib.LayoutFiles(), "templates/post/index.tmpl")
	log.Println("HOME: ", homePageFiles)
	HomePage = &Page{
		Template: template.Must(template.New("home").ParseFiles(homePageFiles...)),
		Layout:   "bootstrap",
	}

	tweetPageFiles := append(lib.LayoutFiles(), "templates/tweet/index.tmpl")
	log.Println("TWEET: ", tweetPageFiles)
	TweetPage = &Page{
		Template: template.Must(template.New("tweet").ParseFiles(tweetPageFiles...)),
		Layout:   "bootstrap",
	}
}
