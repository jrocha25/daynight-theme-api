package handlers

import (
	"net/http"

	"daynight-theme.dev/api/lib"
	"daynight-theme.dev/api/models"
)

func GetAllCountriesHandler(w http.ResponseWriter, r *http.Request) {
	// Get all the countries
	countries, err := lib.GetAllCountries()
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Error fetching countries",
		}
		err = lib.WriteJSONResponse(w, response, http.StatusInternalServerError)
		if err != nil {
			http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
		}
		return
	}

	var returnedCountries []models.APICountry
	for _, country := range countries {
		returnedCountries = append(returnedCountries, models.APICountry{Name: country.Name})
	}

	response := models.APIResponse{
		Success: true,
		Message: "Countries fetched successfully",
		Data:    returnedCountries,
	}

	err = lib.WriteJSONResponse(w, response, http.StatusOK)
	if err != nil {
		http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
	}
}
