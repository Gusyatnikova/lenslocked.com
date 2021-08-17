package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1> My page not found =)</h1>")
}

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h1> Welcome to Natasha's awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a "+
		"href=\"mailto:support@lenlocked.com\">support@lenslocked.com</a>.")
}

func faq (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h2>What is my mood today?</h2>" +
		"<p>Looking forward to a productive day!</p>")
}


func main() {
	r := httprouter.New()
	r.GET("/", home)
	r.GET("/contact", contact)
	r.GET("/faq", faq)
	r.NotFound = http.HandlerFunc(notFoundHandler)
	http.ListenAndServe(":3000", r)
}
