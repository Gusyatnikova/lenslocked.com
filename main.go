package main

import (
	"fmt"
	"net/http"
)

//handlerFunc process incoming web-requests
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welcome to Natasha's awesome site!</h1>" +
		"<h2> Fresh is work now!</h2>")
}

func main() {
	//
	http.HandleFunc("/", handlerFunc)
	// start server and listen on port 3000
	//port is last part of URL, default for HTTP is 80
	//when i type http://www.google.com the browser will use a default port:
	//http://www.google.com:80
	//nil means use a default http handlers
	http.ListenAndServe(":3000", nil)
}