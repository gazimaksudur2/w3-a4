package services

import (
	"encoding/json"
	"os"
	"w3-a4/models"
)


var properties []models.SourceProperty

func init() {
	err := loadProperties()

	if err != nil {
		panic(err)
	}
}

func loadProperties() error {
	file, err := os.ReadFile("data/rental_properties.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &properties)
	if err != nil {
		return  err
	}

	return nil
}

func GetAllProperties() ([]models.SourceProperty, error){
	return properties, nil
}