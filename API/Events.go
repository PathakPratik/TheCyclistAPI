package API

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TheCyclistGoServer/AwsDynamoDb"
)

type Response struct {
	CompletionTime int `json:"CompletionTime"`
	MatchingEvents []struct {
		Categories []struct {
			CategoryID   int    `json:"CategoryID"`
			CategoryName string `json:"CategoryName"`
			Distance     string `json:"Distance"`
			EntryFee     int    `json:"EntryFee"`
			EntryFees    []struct {
				EntryFee int `json:"EntryFee"`
			} `json:"EntryFees"`
			FieldLimit int    `json:"FieldLimit"`
			StartTime  string `json:"StartTime"`
		} `json:"Categories"`
		Distance       int         `json:"Distance"`
		EventAddress   string      `json:"EventAddress"`
		EventCity      string      `json:"EventCity"`
		EventDate      string      `json:"EventDate"`
		EventEndDate   string      `json:"EventEndDate"`
		EventID        int         `json:"EventId"`
		EventName      string      `json:"EventName"`
		EventNotes     string      `json:"EventNotes"`
		EventPermalink string      `json:"EventPermalink"`
		EventState     string      `json:"EventState"`
		EventTypes     []string    `json:"EventTypes"`
		EventURL       string      `json:"EventUrl"`
		EventWebsite   string      `json:"EventWebsite"`
		EventZip       string      `json:"EventZip"`
		Latitude       float64     `json:"Latitude"`
		Longitude      float64     `json:"Longitude"`
		Permit         string      `json:"Permit"`
		PledgeRegURL   interface{} `json:"PledgeRegUrl"`
		PresentedBy    string      `json:"PresentedBy"`
		RegCloseDate   string      `json:"RegCloseDate"`
		RegOpenDate    string      `json:"RegOpenDate"`
	} `json:"MatchingEvents"`
	ResultCount int `json:"ResultCount"`
}

func GetEvents() bool {
	resp, err := http.Get("http://www.BikeReg.com/api/search")
	if err != nil {
		fmt.Println("Error:", err)
	}

	return storeEvents(resp)
}

func storeEvents(resp *http.Response) bool {

	var response Response
	var event AwsDynamoDb.Events

	json.NewDecoder(resp.Body).Decode(&response)

	for _, element := range response.MatchingEvents {
		event.Eventname = element.EventName
		event.EventCity = element.EventCity
		event.EventAddress = element.EventAddress
		event.EventDate = element.EventDate
		event.EventURL = element.EventURL

		AwsDynamoDb.AddEvent(event)
	}

	return true
}
