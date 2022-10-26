package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
)

const alphabhets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// /shortLink - Add a link to the linkList and generate a shorter link
// Example: localhost:9000/shortLink?link=https://google.com
func shortLink(w http.ResponseWriter, r *http.Request) {
	longLink, ok := r.URL.Query()["link"] // Get the URL from the query after ?link=
	if ok {                               // Check if the link is present
		if !validLink(longLink[0]) { //  if the link is NOT valid
			fmt.Fprintf(w, "Could not create shortlink need absolute path link. Ex: /shortLink?link=https://moia.io \n")
			return
		}
		shortLink := createShortLink(5) // Create a random short link, here I am creating a 10 character long link, it can be changed to any length by changing the number
		for _, longLinkinDB := range linkList {
			if longLinkinDB == longLink[0] { // Check if the link already exists in the linkList
				// w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w, "Already have this link http://localhost%s/short/%s \n", portNumber, linkList[longLink[0]]) // If the link already exists, return the short link
				return
			}
		}
		linkList[shortLink] = longLink[0] // Add the short link and the long link to the linkList, useful for checking if the link already exists
		linkList[longLink[0]] = shortLink // Add the long link and the short link to the linkList
		shortenedLink := fmt.Sprintf("Short link for %s is http://localhost%s/short/%s \n", longLink, portNumber, shortLink)

		fmt.Fprint(w, shortenedLink) // Return the short link
		return

	}
	// w.WriteHeader(http.StatusBadRequest) // If the link is not present, return 400
	fmt.Fprintf(w, "Failed to add link \n")
	return
}

// createShortLink - create random short link
func createShortLink(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabhets[rand.Intn(len(alphabhets))] // choosing a random character from the alphabets string
	}
	return string(b) // return the random string
}

// validLink - cvery basic check that the link we're creating is an absolute URL path
func validLink(link string) bool {
	r, err := regexp.Compile("^(http|https)://") // Check that the link starts with http:// or https://
	if err != nil {
		return false
	}
	link = strings.TrimSpace(link)
	// Check if string matches the regex
	return r.MatchString(link)
}
