package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/p3lim/iso8601"
)

// MetaInfo for general meta information for project and api uptime
type MetaInfo struct {
	Uptime  string `json:"uptime"`
	Info    string `json:"info"`
	Version string `json:"version"`
}

var startTime time.Time

// Responds with current API staus
func getAPI(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, MetaInfo{
		Uptime:  iso8601.Format(time.Since(startTime)),
		Info:    "Service for IGC tracks.",
		Version: "v1",
	})
}

// Returns all track IDs if any
func getTrackIDs(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, dbGetTrackIDs())
}

// Adds a new track to db
func postTrack(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
	} else if url, ok := data["url"]; !ok {
		http.Error(w, "Missing url", http.StatusBadRequest)
	} else if id, err := dbPostTrack(url); err != nil {
		http.Error(w, "Url does not contain track data", http.StatusBadRequest)
	} else {
		response := make(map[string]int)
		response["id"] = id
		render.JSON(w, r, response)
	}
}

// Returns track with specific ID if existing
func getTrack(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if data, err := dbGetTrack(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		render.JSON(w, r, data)
	}
}

// Returns specific track field if ID and field exist
func getTrackField(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if data, err := dbGetTrack(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		field := chi.URLParam(r, "field")
		if fieldValue, err := data.dbField(field); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			render.JSON(w, r, fieldValue)
		}
	}
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
	startTime = time.Now()
	dbInit()

	router := chi.NewRouter()
	router.Route("/igcinfo", func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Get("/", getAPI)
			r.Route("/igc", func(r chi.Router) {
				r.Get("/", getTrackIDs)
				r.Post("/", postTrack)
				r.Route("/{id:[0-9]+}", func(r chi.Router) {
					r.Get("/", getTrack)
					r.Get("/{field:[A-Za-z_]+}", getTrackField)
				})
			})
		})
	})
	log.Fatal(http.ListenAndServe(GetPort(), router)) // set listen port
}
