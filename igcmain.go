package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func apiUptime(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	//apiMessage := MetaInfo("Service for IGC tracks.", "v1")
	//b, err := json.Marshal(apiMessage)
	fmt.Fprintf(w, "herokutest")

}

// GetPort port
func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "8080"

	}
	return ":" + port
}
func main() {
	http.HandleFunc("/igcinfo/api", apiUptime)
	err := http.ListenAndServe(GetPort(), nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
