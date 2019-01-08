package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("GO_PORT")
	url := ":" + port
	log.Fatal(http.ListenAndServe(url, nil))
}
