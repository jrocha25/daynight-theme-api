package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"daynight-theme.dev/api/models"
)

const API_URL = "https://api.sunrisesunset.io/json"

func GetDayNightTimes(lat float64, lng float64) (models.DayNightObj, error) {
	query := "?lat=" + fmt.Sprintf("%f", lat) + "&lng=" + fmt.Sprintf("%f", lng) + "&date=today"

	resp, err := http.Get(API_URL + query)
	if err != nil {
		return models.DayNightObj{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.DayNightObj{}, err
	}

	var apiResponse models.SunsetSunriseApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return models.DayNightObj{}, err
	}

	var dayNightObj models.DayNightObj

	dayNightObj.Date = apiResponse.Results.Date
	dayNightObj.Sunrise = apiResponse.Results.Sunrise
	dayNightObj.Sunset = apiResponse.Results.Sunset
	dayNightObj.Timezone = apiResponse.Results.Timezone

	return dayNightObj, nil
}
