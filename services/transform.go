package services

import (
	"encoding/json"
	"w3-a4/models"
)

type Category struct {
	LocationID string `json:"LocationID"`
	Name       string `json:"Name"`
	Type       string `json:"Type"`
}

func parseBreadcrumbs(categoryString string) []string {
	var categories []Category

	err := json.Unmarshal([]byte(categoryString), &categories)

	if err != nil {
		return []string{}
	}
	result := make([]string, 0)

	for _, Item := range categories {
		result = append(result, Item.Name)
	}
	return result
}

func TransformPropety(source models.SourceProperty) models.PropertyResponse {
	response := models.PropertyResponse{}

	response.ID = source.ID
	response.Feed = source.Feed
	response.Published = source.Published

	response.GeoInfo = models.GeoInfo{
		Breadcrumbs: parseBreadcrumbs(source.Categories),
		City:        source.City,
		Country:     source.Country,
		CountryCode: source.CountryCode,
		Name:        source.Display,
		LocationID:  source.LocationID,
		State:       source.State,
		StateAbbr:   source.StateAbbr,
	}
	if len(source.LonLat.Coordinates) >= 2 {
		response.GeoInfo.Lon = source.LonLat.Coordinates[0]
		response.GeoInfo.Lat = source.LonLat.Coordinates[0]
	}
	response.Property = models.PropertyInfo{
		Amenities:    source.AmenityCategories,
		Name:         source.PropertyName,
		Slug:         source.PropertySlug,
		PropertyType: source.PropertyTypeCategory,
		Price:        source.USDPrice,
		ReviewScore:  source.ReviewScoreGeneral,
		StarRating:   source.StarRating,
		Counts: models.PropertyCounts{
			Bathroom:  source.BathroomCount,
			Bedroom:   source.BedroomCount,
			Reviews:   float64(source.NumberOfReview),
			Occupancy: source.Occupancy,
		},
		Image: models.ImageInfo{
			Count:  len(source.Images),
			Images: source.Images,
		},
	}
	return response
}
