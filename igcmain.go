package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlerUptime(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Uptime())
}

func main() {
	http.HandleFunc("/igcinfo/api/", handlerUptime)
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
