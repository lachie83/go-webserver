package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// grab hostname
	hostname, _ := os.Hostname()
	// construction html
	fmt.Fprintf(w, `<html><head></head><body style="text-align: center;">
		<h1><strong>Crocodile Dundee Fan Page</strong></h1>
		<h2><strong>Hostname:</strong> %v<br>
		<strong>Version:</strong> 1.1<br>
		<img src="/static/image2.jpg"><br>
		</h2></body></html>`, hostname)

	// log page hit to stdout
	log.Println("served web page")
}

func main() {
	// log start to stdout
	log.Println("Starting server...")
	// point / at the handler fuction
	http.HandleFunc("/", handler)
	// serve static content from /static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.ListenAndServe(":8080", nil)
}
