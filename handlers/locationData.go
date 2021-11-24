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
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	// Check for correct Content Type
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	// API methods allowed
	allowedMethods := map[string]bool{
		http.MethodPost:    true,
		http.MethodOptions: true,
	}

	if allowedMethods[r.Method] {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Store data in DB
	res := storeLocationData(w, r)

	// Response for Success Case
	if res {
		w.Write([]byte("Success"))
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
