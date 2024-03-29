package main

import (
	. "IMT2681-assignement-2/api"
	"IMT2681-assignement-2/mongodb"
	"IMT2681-assignement-2/webhook"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get port for Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//INIT DATABASE
	mongodb.Global = &mongodb.TracksMongoDB{
		os.Getenv("DB_URL"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_COL_T"),
	}

	//TESTING
	mongodb.G_Webhook = &mongodb.WebhookMongoDB{
		os.Getenv("DB_URL"),
		os.Getenv("DB_NAME"),
		"webhook",
	}

	mongodb.Global.Init()


	// Set up handlers
	r := mux.NewRouter()

	// IGC track handlers
	r.HandleFunc("/paragliding", Redirect).Methods("GET")
	r.HandleFunc("/paragliding/api/", InfoHandler).Methods("GET")
	r.HandleFunc("/paragliding/api/igc", GetAllId).Methods("GET")
	r.HandleFunc("/paragliding/api/igc", AddTrack).Methods("POST")
	r.HandleFunc("/paragliding/api/igc/{id:[0-9]+}", GetTrack).Methods("GET")
	r.HandleFunc("/paragliding/api/igc/{id:[0-9]+}/{prop:[a-z_H]+}", GetTrackProp).Methods("GET")

	// Ticker handlers
	r.HandleFunc("/paragliding/api/ticker/latest", GetLatest).Methods("GET")
	r.HandleFunc("/paragliding/api/ticker/", GetInfoTicker).Methods("GET")
	r.HandleFunc("/paragliding/api/ticker/{time:[0-9]+}", CalcTime).Methods("GET")

	// Webhook handlers:
	r.HandleFunc("/paragliding/api/webhook/new_track/", RegWebH).Methods("POST")
	r.HandleFunc("/paragliding/api/webhook/new_track/{id:[0-9]+}", GetWebH).Methods("GET")
	r.HandleFunc("/paragliding/api/webhook/new_track/{id:[0-9]+}", DelWebH).Methods("DELETE")


	// Admin handlers
	r.HandleFunc("/paragliding/admin/api/tracks_count/{code:[a-z]+}", GetCount).Methods("GET")
	r.HandleFunc("/paragliding/admin/api/tracks/{code:[a-z]+}", DelTracks).Methods("DELETE")

	// Seperate thread
	go webhook.Tracks()

	log.Fatal(http.ListenAndServe(":"+port, r))
}