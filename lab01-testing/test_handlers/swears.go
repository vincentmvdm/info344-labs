package main

import "net/http"

type SwearsHandler struct {
	KnownSwears map[string]bool
}

type SwearCounts map[string]int

func (sh *SwearsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Read through the POST json body and check the text for swear words
	switch r.Method {
	case "POST":

	}
}

func (sh *SwearsHandler) checkSwears(s []string) SwearCounts {
	// make a map so we can count the swears from input
	swearCounts := SwearCounts{}
	// check if each word in s is a swear and
	// increment the counter in swearCounts
	for _, word := range s {
		if sh.KnownSwears[word] {
			swearCounts[word]++
		}
	}

	return swearCounts
}
