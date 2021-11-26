package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/TheCyclistGoServer/AwsDynamoDb"
)

func TripsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	Uid := r.URL.Query().Get("Uid")

	success, trips := AwsDynamoDb.GetItem(Uid)

	if success {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(trips)
	} else {
		http.Error(w, "Invalid Query", http.StatusBadRequest)
	}
}
