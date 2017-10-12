package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestComplimentHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	// e.g. GET /api/projects?page=1&per_page=100
	req, err := http.NewRequest("GET", "/compliments", nil)
	if err != nil {
		t.Fatal(err)
	}

	// modify the url query to add the name query parameters
	name := "ethan"
	q := req.URL.Query()
	q.Add("name", name)
	req.URL.RawQuery = q.Encode() // re encode the query string

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ComplimentHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expectedPrefix := fmt.Sprintf("%s, you are", name)
	if !strings.HasPrefix(rr.Body.String(), expectedPrefix) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedPrefix)
	}
}
