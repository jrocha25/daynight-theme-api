package models

type DayNightObj struct {
	Date     string `json:"date"`
	Sunrise  string `json:"sunrise"`
	Sunset   string `json:"sunset"`
	Timezone string `json:"timezone"`
}
