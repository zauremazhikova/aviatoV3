package repositories

import (
	"aviatoV3/internal/database"
	"aviatoV3/internal/entities"
	"fmt"
	"time"
)

func GetCities() (a []*entities.City, err error) {
	cities := make([]*entities.City, 0)

	db := database.DB()
	rows, err := db.Query("SELECT ID, NAME, COUNTRY_ID, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM cities")
	_ = db.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var city entities.City
		var countryID string
		err := rows.Scan(&city.ID, &city.Name, &countryID, &city.CreatedAt, &city.UpdatedAt, &city.DeletedAt)
		if err != nil {
			fmt.Println(err)
			return cities, err
		} else {
			currentCountry, _ := GetCountry(countryID)
			city.Country = *currentCountry
			cities = append(cities, &city)
		}
	}

	return cities, nil
}

func GetCity(id string) (*entities.City, error) {

	db := database.DB()
	rows, err := db.Query("SELECT ID, NAME, COUNTRY_ID, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM cities WHERE ID = $1", id)
	_ = db.Close()
	if err != nil {
		return nil, err
	}

	var city entities.City
	var countryID string
	for rows.Next() {
		err := rows.Scan(&city.ID, &city.Name, &countryID, &city.CreatedAt, &city.UpdatedAt, &city.DeletedAt)
		if err != nil {
			return &entities.City{}, err
		}
	}
	currentCountry, _ := GetCountry(countryID)
	city.Country = *currentCountry

	return &city, nil
}

func CreateCity(city *entities.City) error {
	db := database.DB()
	_, err := db.Query("INSERT INTO cities (name, country_id, created_at) VALUES ($1, $2, $3)", city.Name, city.Country.ID, time.Now())

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func UpdateCity(city *entities.City) error {
	db := database.DB()
	_, err := db.Query("UPDATE cities SET name = $2, updated_at = $3 WHERE id = $1", city.ID, city.Name, time.Now())

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func DeleteCity(id string) error {
	db := database.DB()
	_, err := db.Query("UPDATE cities SET deleted_at = $1 WHERE id = $2", time.Now(), id)

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}
