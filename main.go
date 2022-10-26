package main

import (
	"log"
	"net/http"
)

var (
	linkList   map[string]string
	portNumber = ":9000"
)

func main() {
	linkList = map[string]string{}

	http.HandleFunc("/shortLink", shortLink)   // /short endpoint to shorten the link
	http.HandleFunc("/short/", visitShortLink) // /short/{shortLink}  to visit the link
	http.HandleFunc("/", home)                 // Home page that can be visited by localhost:9000, check Readme on how to change portNumber
	log.Println("Server started on port", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}
