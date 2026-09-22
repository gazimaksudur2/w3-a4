package services

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"w3-a4/models"
)


var sourceProperties []models.SourceProperty
var ErrPropertyNotFound = errors.New("Property not found")

var once sync.Once
var loadError error

func loadProperties() error {
	file, err := os.ReadFile("data/rental_properties.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &sourceProperties)

	return err
}

func GetAllProperties() ([]models.SourceProperty, error){
	once.Do(func(){
		loadError = loadProperties()
	})

	if loadError != nil {
		return nil, loadError
	}

	return sourceProperties, nil
}

func GetPropertyByID(id string) (*models.PropertyResponse, error) {
	properties, err := GetAllProperties()
	if err!=nil {
		return nil, err
	}

	for _, property := range properties {
		if property.ID == id{
			response := TransformProperty(property)
			return &response, nil
		}
	}
	return nil, ErrPropertyNotFound
}

func ListProperties(
	filter models.PropertyFilter,
) (models.PropertyListResponse, error) {
	sourceProperties, err := GetAllProperties()

	if err != nil {
		return models.PropertyListResponse{}, err
	}

	filtered := FilterProperties(sourceProperties, filter)
	items := make([]models.PropertyResponse, 0)

	for _, property := range filtered {
		items = append(items, TransformProperty(property))
	}
	if filter.Limit > 0 && len(items)>filter.Limit {
		items = items[:filter.Limit]
	}
	return models.PropertyListResponse{
		Result: models.Result{
			Count: len(items),
			Items: items,
		},
	}, nil
}