package main

import (
	"fmt"
	"log"
	"net/http"
)

func apiUptime(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	//apiMessage := MetaInfo("Service for IGC tracks.", "v1")
	//b, err := json.Marshal(apiMessage)
	fmt.Fprintf(w, "herokutest")

}

func main() {
	http.HandleFunc("/igcinfo/api", apiUptime)
	err := http.ListenAndServe(":", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
