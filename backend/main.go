package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Backend running on port :9001")

	http.ListenAndServe(":9001", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")

	body, _ := io.ReadAll(r.Body)

	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Query: %s\n", r.URL.RawQuery)
	fmt.Fprintf(w, "Body: %s\n", string(body))
	fmt.Fprintf(w, "User-Agent: %s\n", r.UserAgent())
	fmt.Println(r.Header.Get("X-Forwarded-By"))
}
