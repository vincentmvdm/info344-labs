package main

import "net/http"

type ComplimentHandler struct {
	Compliments []string
}

func (ch *ComplimentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
