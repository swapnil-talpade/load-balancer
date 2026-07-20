package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "9001", "backend port")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Response from backend %s\n", *port)
	})

	log.Printf("Backend running on :%s", *port)

	log.Fatal(
		http.ListenAndServe(":"+*port, nil),
	)
}
