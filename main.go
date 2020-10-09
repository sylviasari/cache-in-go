package main

import(
	"log"
	"net/http"
)

func main() {
	initCache()
	handleRequests()
	log.Fatal(http.ListenAndServe(":15000", nil))
}
