package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

// MetaInfo struct for api uptime and meta info
type MetaInfo struct {
	Uptime  string `json: "uptime"`
	Info    string `json: "info"`
	Version string `json: "version"`
}

// Uptime returns app uptime in the ISO 8601 standard
func Uptime() MetaInfo {
	return MetaInfo{"P", "info", "v"}
}

func meta(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	json.NewEncoder(w).Encode(Uptime())
}

func trackIds(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func trackRegistration(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func metaID(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func metaIDSingle(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

// GetPort port
func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "8080"

	}
	return ":" + port
}
func noFunction(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {
	router := chi.NewRouter()
	router.Route("/igcinfo", func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Get("/", meta)
			r.Route("/igc", func(r chi.Router) {
				r.Get("/", trackIds)
				r.Post("/", trackRegistration)
				r.Route("/{id:[0-9]+}", func(r chi.Router) {
					r.Get("/", metaID)
					r.Get("/{field:[A-Za-z_]+}", metaIDSingle)
				})
			})
		})
	})
	http.ListenAndServe(GetPort(), router) // set listen port

}
