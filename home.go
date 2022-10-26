package main

import (
	"fmt"
	"net/http"
)

// Home - Home http request
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	var response string
	for shortLink, longLink := range linkList {
		if len(longLink) != 5 { // Check if the link is a long link
			response += fmt.Sprintf("<br>ShortLink: <a href=\"http://localhost%s/short/%s\">http://localhost%s/short/%s</a> \t\t LongLink: <a href=\"%s\">%s</a>", portNumber, shortLink, portNumber, shortLink, longLink, longLink)
		}
	}
	fmt.Fprintf(w, "<h2>URL Shortener!<h2><br>\n")
	fmt.Fprint(w, response)
	return
}
