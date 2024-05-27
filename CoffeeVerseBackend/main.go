package main

import (
	"CoffeeVerseBackend/BrewSetting"
	"encoding/json"
	"fmt"
	"net/http"
)

func insertBrewSetting(w http.ResponseWriter, req *http.Request) {
	var receivedBrewSetting BrewSetting.BrewSetting
	err := json.NewDecoder(req.Body).Decode(&receivedBrewSetting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(receivedBrewSetting)
}

func main() {
	http.HandleFunc("/insertBrewSetting", insertBrewSetting)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
