package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type hdump struct{}

func (h hdump) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(requestDump))
}
