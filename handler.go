package main

import (
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/github", github)
}
