package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello World!"))
	})

	s := http.Server{
		Addr:    ":443",
		Handler: mux,
		TLSConfig: &tls.Config{
			NextProtos: []string{"h2", "http/1.1"},
		},
	}

	fmt.Printf("Server listening on %s\n", s.Addr)
	log.Fatal(s.ListenAndServeTLS("certs/localhost.crt", "certs/localhost.key"))
}
