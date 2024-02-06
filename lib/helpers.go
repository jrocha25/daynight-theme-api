package lib

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"daynight-theme.dev/api/models"
)

type CSVRecord struct {
	Country   string
	Latitude  string
	Longitude string
}

func WriteJSONResponse(w http.ResponseWriter, response models.APIResponse, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}

	return nil
}

func getCountriesInfo() ([]CSVRecord, error) {
	const url = "https://gist.githubusercontent.com/jrocha25/2a49aae90e7f014136211e5536ae3138/raw/355eb56e164ddc3cd1a9467c524422cb674e71a9/country-capital-lat-long-population.csv"

	// Get the CSV file
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Create a new CSV reader reading from the response body
	reader := csv.NewReader(resp.Body)

	// Skip the first line (header)
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	var records []CSVRecord

	// Iterate through the remaining records
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Extract the desired fields
		country := record[0]
		latitude := record[2]
		longitude := record[3]

		// Append to the records slice
		records = append(records, CSVRecord{
			Country:   country,
			Latitude:  latitude,
			Longitude: longitude,
		})
	}

	return records, nil
}

func SeedDatabase() error {
	records, err := getCountriesInfo()
	if err != nil {
		return err
	}

	for _, record := range records {
		_, err := dbClient.Exec("INSERT INTO countries (name, latitude, longitude) SELECT ?, ?, ? WHERE NOT EXISTS (SELECT 1 FROM countries WHERE name = ?)", record.Country, record.Latitude, record.Longitude, record.Country)
		if err != nil {
			return err
		}
	}

	return nil
}
