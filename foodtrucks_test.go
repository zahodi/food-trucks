package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type FoodTruck struct {
	LocationID     string `json:"locationid"`
	Applicant      string `json:"applicant"`
	FacilityType   string `json:"facilitytype"`
	LocationDesc   string `json:"locationdesc"`
	Address        string `json:"address"`
	Permit         string `json:"permit"`
	Status         string `json:"status"`
	FoodItems      string `json:"fooditems"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	Schedule       string `json:"schedule"`
	DaysHours      string `json:"dayshours"`
	NOISent        string `json:"noisent"`
	Approved       string `json:"approved"`
	Received       string `json:"received"`
	PriorPermit    string `json:"priorpermit"`
	ExpirationDate string `json:"expirationdate"`
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/foodtrucks", getFoodTrucks).Methods("GET")
	r.HandleFunc("/foodtrucks/search", searchFoodTrucks).Methods("GET")
	return r
}

func TestGetFoodTrucks(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/foodtrucks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []FoodTruck
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, response)
}

func TestSearchFoodTrucks(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/foodtrucks/search?food=burger", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, response)
}

func TestSearchFoodTrucksNoResults(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/foodtrucks/search?food=nonexistentfood", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Empty(t, response)
}

func init() {
	// Initialize the foodTrucks variable with some sample data for testing purposes
	foodTrucks = []map[string]string{
		{
			"locationid":     "1",
			"applicant":      "Food Truck A",
			"facilitytype":   "Truck",
			"locationdesc":   "Desc 1",
			"address":        "123 Fake St",
			"permit":         "Permit A",
			"status":         "ACTIVE",
			"fooditems":      "Burgers",
			"latitude":       "37.78",
			"longitude":      "-122.41",
			"schedule":       "https://schedule.com",
			"dayshours":      "Mon-Fri 10-2",
			"noisent":        "",
			"approved":       "2021-01-01",
			"received":       "2020-12-01",
			"priorpermit":    "0",
			"expirationdate": "2022-01-01",
		},
		{
			"locationid":     "2",
			"applicant":      "Food Truck B",
			"facilitytype":   "Truck",
			"locationdesc":   "Desc 2",
			"address":        "456 Fake St",
			"permit":         "Permit B",
			"status":         "ACTIVE",
			"fooditems":      "Pizza",
			"latitude":       "37.79",
			"longitude":      "-122.42",
			"schedule":       "https://schedule.com",
			"dayshours":      "Mon-Fri 10-2",
			"noisent":        "",
			"approved":       "2021-01-01",
			"received":       "2020-12-01",
			"priorpermit":    "0",
			"expirationdate": "2022-01-01",
		},
	}
}
