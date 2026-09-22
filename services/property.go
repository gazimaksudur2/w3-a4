package services

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"w3-a4/models"
)


var properties []models.SourceProperty

var once sync.Once
var loadError error

func loadProperties() error {
	file, err := os.ReadFile("data/rental_properties.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &properties)

	return err
}

func GetAllProperties() ([]models.SourceProperty, error){
	once.Do(func(){
		loadError = loadProperties()
	})

	if loadError != nil {
		return nil, loadError
	}

	return properties, nil
}

func GetPropertyByID(id string) (*models.PropertyResponse, error) {
	properties, err := GetAllProperties()
	if err!=nil {
		return nil, err
	}

	for _, property := range properties {
		if property.ID == id{
			response := TransformPropety(property)
			return &response, nil
		}
	}
	return nil, errors.New("property not found")
}