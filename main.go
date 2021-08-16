package main

import (
	"fmt"
	"net/http"
)

//handlerFunc process incoming web-requests
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1> Welcome to Natasha's awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a "+
			"href=\"mailto:support@lenlocked.com\">support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h2>We could not find the page you were looking for :(</h2>" +
			"<p>Plese email us if you keep being sent to an invalid page.</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
