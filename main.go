package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func github(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.github.com/status")
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.NewEncoder(w).Encode(string(data))
}

func handleRequests() {
	http.HandleFunc("/github", github)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
