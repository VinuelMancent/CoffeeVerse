package main

import (
	"CoffeeVerseBackend/BrewSetting"
	"encoding/json"
	"fmt"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Referrer-Policy", "no-referrer")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func insertBrewSetting(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request incoming from client")
	enableCors(&w)
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
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
