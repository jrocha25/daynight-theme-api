package lib

import (
	"encoding/json"
	"net/http"

	"daynight-theme.dev/api/models"
)

func WriteJSONResponse(w http.ResponseWriter, response models.APIResponse, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}

	return nil
}
