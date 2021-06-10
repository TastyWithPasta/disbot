package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

// TODO: Create basic handler to dump REST requests to console.
// TODO: Test handler by using testbot1 and dumping request

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, string(requestDump))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
