package handlers

import (
	"log"
	"net/http"

	"daynight-theme.dev/api/lib"
	"daynight-theme.dev/api/models"
)

func GetTimesHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")

	log.Println("Fetching day/night times for", location)

	// Get the country info
	countryInfo, err := lib.GetCountryInfo(location)
	if err != nil {
		log.Println("Error fetching country info:", err)
		response := models.APIResponse{
			Success: false,
			Message: "Error fetching country info",
		}
		err = lib.WriteJSONResponse(w, response, http.StatusInternalServerError)
		if err != nil {
			http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
		}
		return
	}

	returnedTimes, err := lib.GetDayNightTimes(countryInfo.Latitude, countryInfo.Longitude)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Error fetching day/night times",
		}
		err = lib.WriteJSONResponse(w, response, http.StatusInternalServerError)
		if err != nil {
			http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
		}
		return
	}

	response := models.APIResponse{
		Success: true,
		Message: "Day/Night times fetched successfully",
		Data:    returnedTimes,
	}

	err = lib.WriteJSONResponse(w, response, http.StatusOK)
	if err != nil {
		http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
	}
}
