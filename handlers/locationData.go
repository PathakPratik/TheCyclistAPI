package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/golang/gddo/httputil/header"

	"github.com/TheCyclistGoServer/AwsDynamoDb"
)

func LocationDataHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

	// Check for correct Content Type
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Store data in DB
	res := storeLocationData(w, r)

	// Response for Success Case
	if res {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
		return
	}

}

func storeLocationData(w http.ResponseWriter, r *http.Request) bool {

	var trackingData AwsDynamoDb.TrackingData

	json.NewDecoder(r.Body).Decode(&trackingData)

	result, err := govalidator.ValidateStruct(trackingData)

	if !result {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	AwsDynamoDb.AddItem(trackingData)

	return true
}
