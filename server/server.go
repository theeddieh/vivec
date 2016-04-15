package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func info(w http.ResponseWriter, r *http.Request) {
	content := fmt.Sprint("Method: ", r.Method, "\n",
		"URL: ", r.URL, "\n",
		"URI: ", r.RequestURI, "\n",
		"Protocol: ", r.Proto, "\n",
		"Header: ", r.Header, "\n",
		"Body: ", r.Body)
	io.WriteString(w, content)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/info", info)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
