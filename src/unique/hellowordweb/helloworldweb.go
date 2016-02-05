package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World %s!", r.URL.Path[1:])
}

func main() {
	p := "9080"
	if len(os.Args) > 1 {
		p = os.Args[1]
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+p, nil)
}
