package api

import (
	"IMT2681-assignement-2/mongodb"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

/*
Delete every track in db
*/
func DelTracks(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	code := vars["code"]

	if code == "admin" {
		ok := mongodb.Global.DelAll()
		if !ok {
			http.Error(w, "Failed to delete", http.StatusNotFound)
			return
		}
		fmt.Fprintf(w,"Deleted "+strconv.Itoa(mongodb.Global.Count()))
	} else {
		http.Error(w, "No access", http.StatusForbidden)
	}
}

/*
Get number of tracks in db
*/
func GetCount(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	code := vars["code"]

	if code == "admin" {
		allTrack := mongodb.Global.Count()
		fmt.Println(allTrack)
		fmt.Fprintf(w,strconv.Itoa(allTrack))
	} else {
		http.Error(w,"No access",http.StatusForbidden)
	}
}