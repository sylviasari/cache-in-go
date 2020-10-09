package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)

func github(w http.ResponseWriter, r *http.Request) {
	err, result := getCache("github")
	fmt.Printf("\n err %+v result %+v", err, result)

	if err != nil{
		resp, err := http.Get("https://api.github.com/status")
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		data, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		result = string(data)

		setCache("github", result)
	}

	json.NewEncoder(w).Encode(result)
}
