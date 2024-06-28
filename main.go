package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var foodTrucks []map[string]string

func loadCSVFromURL(url string) ([]map[string]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch CSV: %v", resp.Status)
	}

	reader := csv.NewReader(resp.Body)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("CSV file is empty")
	}

	var headers []string
	// convert headers to lowercase
	for _, i := range records[0] {
		headers = append(headers, strings.ToLower(i))
	}

	var foodTrucks []map[string]string
	for _, record := range records[1:] {
		foodTruck := make(map[string]string)
		for i, header := range headers {
			foodTruck[header] = record[i]
		}
		foodTrucks = append(foodTrucks, foodTruck)
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

	uniqueApplicants := make(map[string]bool)
	var result []string

	for _, truck := range foodTrucks {
		foodItems := truck["fooditems"]
		if strings.Contains(strings.ToLower(foodItems), strings.ToLower(query)) {
			applicant := truck["applicant"]
			if _, exists := uniqueApplicants[applicant]; !exists {
				uniqueApplicants[applicant] = true
				result = append(result, truck["applicant"])
			}
		}
	}

	if len(result) == 0 {
		fmt.Println("No results found for query:", query)
	} else {
		fmt.Println("Results found for query:", query)
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	url := "https://data.sfgov.org/api/views/rqzj-sfat/rows.csv"
	var err error
	foodTrucks, err = loadCSVFromURL(url)
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
