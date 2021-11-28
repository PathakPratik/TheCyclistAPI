package handlers

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/TheCyclistGoServer/AwsDynamoDb"
)

func decorator(trips []AwsDynamoDb.TrackingData) [][]AwsDynamoDb.TrackingData {
	sort.Slice(trips, func(i, j int) bool {
		return trips[i].Timestamp > trips[j].Timestamp
	})

	var tripPoints [][]AwsDynamoDb.TrackingData
	var res []AwsDynamoDb.TrackingData
	currTrip := trips[0].TripId
	for _, trip := range trips {
		if trip.TripId != currTrip {
			tripPoints = append(tripPoints, res)
			res = []AwsDynamoDb.TrackingData{}
			currTrip = trip.TripId
		}

		res = append(res, AwsDynamoDb.TrackingData{
			TripId: trip.TripId, UserId: trip.UserId, Timer: trip.Timer, RecordId: trip.RecordId, Latitude: trip.Latitude, Longitude: trip.Longitude, Timestamp: trip.Timestamp,
		})
	}
	tripPoints = append(tripPoints, res)

	return tripPoints
}

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
		json.NewEncoder(w).Encode(decorator(trips))
	} else {
		http.Error(w, "Invalid Query", http.StatusBadRequest)
	}
}
