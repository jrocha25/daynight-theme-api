package handlers

import (
	"net/http"

	"daynight-theme.dev/api/lib"
	"daynight-theme.dev/api/models"
)

func GetTimesHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// location := vars["location"]

	//TODO: Verify location is valid

	lat := 38.907192
	lng := -77.036873

	returnedTimes, err := lib.GetDayNightTimes(lat, lng)
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
