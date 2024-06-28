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

var foodTrucks []map[string]string

func loadCSV(filePath string) ([]map[string]string, error) {
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

	if len(records) == 0 {
		return nil, fmt.Errorf("CSV file is empty")
	}

	headers := records[0]
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
	// var result []map[string]string
	var result []string

	for _, truck := range foodTrucks {
		foodItems := truck["FoodItems"] // or use the dynamic column name if it's different
		if strings.Contains(strings.ToLower(foodItems), strings.ToLower(query)) {
			applicant := truck["Applicant"] // or use the dynamic column name if it's different
			if _, exists := uniqueApplicants[applicant]; !exists {
				uniqueApplicants[applicant] = true
				result = append(result, truck["Applicant"])
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
