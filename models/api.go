package models

type SunsetSunriseApiResponse struct {
	Results struct {
		Date       string `json:"date"`
		Sunrise    string `json:"sunrise"`
		Sunset     string `json:"sunset"`
		FirstLight string `json:"first_light"`
		LastLight  string `json:"last_light"`
		Dawn       string `json:"dawn"`
		Dusk       string `json:"dusk"`
		SolarNoon  string `json:"solar_noon"`
		GoldenHour string `json:"golden_hour"`
		DayLength  string `json:"day_length"`
		Timezone   string `json:"timezone"`
		UtcOffset  int    `json:"utc_offset"`
	} `json:"results"`
	Status string `json:"status"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type APICountry struct {
	Name string `json:"name"`
}
