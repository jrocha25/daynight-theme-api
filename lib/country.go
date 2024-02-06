package lib

import "daynight-theme.dev/api/models"

func GetCountryInfo(name string) (models.Country, error) {
	var countryInfo models.Country

	// Get the country info from the database
	err := dbClient.QueryRow("SELECT name, latitude, longitude FROM countries WHERE name = ?", name).Scan(&countryInfo.Name, &countryInfo.Latitude, &countryInfo.Longitude)
	if err != nil {
		return models.Country{}, err
	}

	return countryInfo, nil
}
