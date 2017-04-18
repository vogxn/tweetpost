package views

import (
	"html/template"
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
	layoutFiles := lib.LayoutFiles()

	homePageFiles := append(layoutFiles, "templates/post/index.tmpl")
	HomePage = &Page{
		Template: template.Must(template.New("home").ParseFiles(homePageFiles...)),
		Layout:   "bootstrap",
	}

	tweetPageFiles := append(layoutFiles, "templates/tweet/index.tmpl")
	TweetPage = &Page{
		Template: template.Must(template.New("tweet").ParseFiles(tweetPageFiles...)),
		Layout:   "bootstrap",
	}
}
