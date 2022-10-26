package main

import (
	"fmt"
	"net/http"
	"strings"
)

// visitShortLink - Find link that matches the shortened link in the linkList
func visitShortLink(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path                   // Get the path from the URL
	pathArgs := strings.Split(path, "/") // Split the path into a slice
	if len(pathArgs[2]) < 1 {
		w.WriteHeader(http.StatusNotFound)                            // If the path is not long enough, return 404
		requestURL := fmt.Sprintf("http://localhost%s", portNumber)   // Create the URL to redirect
		http.Redirect(w, r, requestURL, http.StatusTemporaryRedirect) // Redirect
		return
	}
	http.Redirect(w, r, linkList[pathArgs[2]], http.StatusTemporaryRedirect) // Redirect to the actual link
	return
}
