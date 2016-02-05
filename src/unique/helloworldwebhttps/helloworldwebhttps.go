package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World %s!", r.URL.Path[1:])
}

func main() {
	p := "9080"
	ps := "9443"
	if len(os.Args) > 1 {
		p = os.Args[1]
	}
	if len(os.Args) > 2 {
		ps = os.Args[2]
	}

	http.HandleFunc("/", handler)

	// Start HTTP server on port 80
	go func() {
		err := http.ListenAndServe(":"+p, nil)
		if err != nil {
			log.Fatal("ListenAndServe "+p+": ", err)
		}
	}()

	CA_Pool := x509.NewCertPool()

	pemData, err := ioutil.ReadFile("localhost.crt")
	if err != nil {
		log.Fatal("localhost.crt unavailable: ", err)
	}
	CA_Pool.AppendCertsFromPEM(pemData)
	config := &tls.Config{RootCAs: CA_Pool}
	server := &http.Server{Addr: ":" + ps, TLSConfig: config}

	// Start HTTP server on port ps (9443 default)
	err = server.ListenAndServeTLS("localhost.crt", "localhost.key")
	if err != nil {
		log.Fatal("ListenAndServe "+ps+": ", err)
	}
}
