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

func GetAllCountries() ([]models.Country, error) {
	var countries []models.Country

	// Get all the countries from the database
	rows, err := dbClient.Query("SELECT name FROM countries")
	if err != nil {
		return []models.Country{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var country models.Country
		err = rows.Scan(&country.Name)
		if err != nil {
			return []models.Country{}, err
		}
		countries = append(countries, country)
	}

	return countries, nil
}
