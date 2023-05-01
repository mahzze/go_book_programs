package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler echoes the HTTP request. This time it is used as a kind of tool to debug.
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "%s %s %s\n", request.Method, request.URL, request.Proto)
	for key, value := range request.Header {
		fmt.Fprintf(writer, "Header[%q] = %q\n", key, value)
	}
	fmt.Fprintf(writer, "Host = %q\n", request.Host)
	fmt.Fprintf(writer, "Remote Address= %q\n", request.RemoteAddr)

	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}

	for key, value := range request.Form {
		fmt.Fprintf(writer, "Form[%q] = %q\n", key, value)
	}
}
