package handlers

import "net/http"

type SwearsHandler struct {
	KnownSwears []string
}

func (sh *SwearsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Read through the POST json body and check the text for swear words
	switch r.Method {
	case "POST":

	}
}

func (sh *SwearsHandler) checkSwears() {

}
