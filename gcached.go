package main

import (
	"fmt"
	"log"
	"net/http"
)

// the global buckets list
// var buckets map[string]Bucket
var buckets = map[string]Bucket{}

func main() {

	// the global buckets list
	buckets = map[string]Bucket{}

	fmt.Println("gcached starting up... Try :8080")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
