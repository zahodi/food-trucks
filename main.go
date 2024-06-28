package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type FoodTruck struct {
	LocationID     string `json:"locationid"`
	Applicant      string `json:"applicant"`
	FacilityType   string `json:"facilitytype"`
	Cnn            string `json:"cnn"`
	LocationDesc   string `json:"locationdesc"`
	Address        string `json:"address"`
	BlockLot       string `json:"blocklot"`
	Block          string `json:"block"`
	Lot            string `json:"lot"`
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

var foodTrucks []FoodTruck

func loadCSV(filePath string) ([]FoodTruck, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var foodTrucks []FoodTruck
	for _, record := range records[1:] {
		// Print FoodItems column for debugging
		fmt.Println("FoodItems:", record[7])

		foodTrucks = append(foodTrucks, FoodTruck{
			LocationID:     record[0],
			Applicant:      record[1],
			FacilityType:   record[2],
			Cnn:            record[3],
			LocationDesc:   record[4],
			Address:        record[5],
			BlockLot:       record[6],
			Block:          record[7],
			Lot:            record[8],
			Permit:         record[6],
			Status:         record[7],
			NOISent:        record[8],
			Latitude:       record[9],
			Longitude:      record[10],
			Schedule:       record[11],
			FoodItems:      record[11],
			DaysHours:      record[12],
			Approved:       record[13],
			Received:       record[14],
			PriorPermit:    record[15],
			ExpirationDate: record[16],
		})
	}

	return foodTrucks, nil
}

func getFoodTrucks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foodTrucks)
}

func searchFoodTrucks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query().Get("food")
	if query == "" {
		http.Error(w, "Query parameter 'food' is required", http.StatusBadRequest)
		return
	}

	var result []FoodTruck
	for _, truck := range foodTrucks {
		if strings.Contains(strings.ToLower(truck.FoodItems), strings.ToLower(query)) {
			result = append(result, truck)
		}
	}

	if len(result) == 0 {
		fmt.Println("No results found for query:", query)
		fmt.Println(foodTrucks[1])
	} else {
		fmt.Println("Results found for query:", query)
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	var err error
	foodTrucks, err = loadCSV("rows.csv")
	if err != nil {
		fmt.Printf("Error loading CSV file: %v\n", err)
		return
	}

	fmt.Println("Loaded", len(foodTrucks), "food trucks")

	r := mux.NewRouter()
	r.HandleFunc("/foodtrucks", getFoodTrucks).Methods("GET")
	r.HandleFunc("/foodtrucks/search", searchFoodTrucks).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}
