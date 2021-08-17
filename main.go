package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h1> Welcome to Natasha's awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a "+
		"href=\"mailto:support@lenlocked.com\">support@lenslocked.com</a>.")
}

func faq (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")
	fmt.Fprint(w, "<h2>What is my mood today?</h2>" +
		"<p>Looking forward to a productive day!</p>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}
