package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Static maps static files
func Static(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Disable listing directories
	if strings.HasSuffix(r.URL.Path, "/") {
		return
	}
	log.Println("serving static file", r.URL.Path)
	http.ServeFile(w, r, r.URL.Path[1:])
}
