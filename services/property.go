package services

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"w3-a4/models"

	"github.com/beego/beego/v2/core/logs"
)


var sourceProperties []models.SourceProperty
var ErrPropertyNotFound = errors.New("property not found")

var once sync.Once
var loadError error

func loadProperties() error {
	_, currentFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(currentFile)
	path := filepath.Join(
	baseDir,
	"../data/rental_properties.json",
)

file, err := os.ReadFile(path)
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
		logs.Error("failed loading property data %v", err)
		return nil, err
	}

	return FindPropertyByID(properties, id)
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

func FindPropertyByID(properties []models.SourceProperty, id string) (*models.PropertyResponse, error) {

	for _, property := range properties {

		if property.ID == id {

			response := TransformProperty(property)

			return &response, nil
		}
	}

	return nil, ErrPropertyNotFound
}