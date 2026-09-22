package services

import (
	"encoding/json"
	"os"
	"sync"
	"w3-a4/models"
)


var properties []models.SourceProperty

var once sync.Once
var loadError error

func GetAllProperties() ([]models.SourceProperty, error){
	once.Do(func(){
		loadError = loadProperties()
	})

	if loadError != nil {
		return nil, loadError
	}

	return properties, nil
}

func loadProperties() error {
	file, err := os.ReadFile("data/rental_properties.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &properties)

	return err
}