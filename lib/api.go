package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"daynight-theme.dev/api/models"
)

const API_URL = "https://api.sunrisesunset.io/json"

func GetDayNightTimes(lat string, lng string) (models.DayNightObj, error) {
	query := "?lat=" + fmt.Sprintf("%s", lat) + "&lng=" + fmt.Sprintf("%s", lng) + "&date=today"

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

	// Parse the sunrise time and format it in 24-hour format
	sunrise, err := time.Parse("3:04:05 PM", apiResponse.Results.Sunrise)
	if err != nil {
		return models.DayNightObj{}, err
	}
	dayNightObj.Sunrise = sunrise.Format("15:04:05")

	// Parse the sunset time and format it in 24-hour format
	sunset, err := time.Parse("3:04:05 PM", apiResponse.Results.Sunset)
	if err != nil {
		return models.DayNightObj{}, err
	}
	dayNightObj.Sunset = sunset.Format("15:04:05")

	dayNightObj.Timezone = apiResponse.Results.Timezone

	return dayNightObj, nil
}
