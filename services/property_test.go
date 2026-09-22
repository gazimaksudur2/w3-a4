package services

import (
	"testing"
	"w3-a4/models"
)

func createTestProperties() []models.SourceProperty {
	return []models.SourceProperty{
		{
			ID: "1",
			Feed: 11,
			Published: false,
			USDPrice: 100,
			StarRating: 5,
			ReviewScoreGeneral: 4.5,
			NumberOfReview: 50,
			PropertyTypeCategory: "Hotel",
			BedroomCount: 3,
			AmenityCategories: []string{
				"Pool",
				"Internet",
			},
		},
		{
			ID: "2",
			Feed: 12,
			Published: true,
			USDPrice: 200,
			StarRating: 3,
			ReviewScoreGeneral: 3.5,
			NumberOfReview: 10,
			PropertyTypeCategory: "Apartment",
			BedroomCount: 1,
			AmenityCategories: []string{
				"Parking",
			},
		},
		{
			ID: "3",
			Feed: 11,
			Published: true,
			USDPrice: 300,
			StarRating: 4,
			ReviewScoreGeneral: 4.8,
			NumberOfReview: 100,
			PropertyTypeCategory: "Villa",
			BedroomCount: 4,
			AmenityCategories: []string{
				"Gym",
			},
		},
	}
}

func TestTransformProperty(t *testing.T) {
	source := models.SourceProperty{
		ID: "TEST-1",
		Images: []string{
			"a.jpg",
			"b.jpg",
		},
		LonLat: models.LonLat{
			Coordinates: []float64{
				100,
				20,
			},
		},
		Categories: `[
			{
				"Name": "Japan"
			},
			{
				"Name": "Tokyo"
			}
		]`,
	}

	result := TransformProperty(source)

	if result.ID != "TEST-1" {
		t.Error("ID mapping failed")
	}
	if result.Property.Image.Count != 2 {
		t.Error("Image count failed")
	}
	if len(result.GeoInfo.Breadcrumbs) != 2 {
		t.Error("Breadcrumb parsing failed")
	}
}

func TestFilterProperties_AND(t *testing.T) {
	properties := createTestProperties()

	published := false
	feed := 11

	filter := models.PropertyFilter{
		Feed: &feed,
		Published: &published,
	}
	
	result := FilterProperties(properties, filter)
	if len(result)!=1 {
		t.Errorf(
			"expected 1 result got %d", len(result),
		)
	}
	if result[0].ID!="1" {
		t.Errorf(
			"expected ID 1 got %s", result[0].ID,
		)
	}
}

func TestFilterProperties_PriceRange(t *testing.T){
	properties := createTestProperties()
	min := 50.0
	max := 150.0

	filter := models.PropertyFilter{
		MinPrice: &min,
		MaxPrice: &max,
	}
	result := FilterProperties(properties, filter)
	if len(result) != 1 {
		t.Errorf("expected 1 result got %d", len(result))
	}
}

func TestFilterProperties_AmenitiesOR(t *testing.T) {
	properties := createTestProperties()
	filter := models.PropertyFilter{
		Amenities: []string{
			"Internet",
			"Parking",
		},
	}
	result := FilterProperties(properties, filter)
	if len(result) != 2 {
		t.Errorf("expected 2 results got %d", len(result))
	}
}

func TestFilterProperties_Combined(t *testing.T) {
	properties := createTestProperties()
	feed := 11
	filter := models.PropertyFilter{
		Feed: &feed,
		Amenities: []string{
			"Internet",
		},
	}
	result := FilterProperties(properties, filter)
	if len(result)!=1 {
		t.Errorf("expected 1 result got %d", len(result))
	}
	if result[0].ID != "1" {
		t.Errorf("wrong property returned")
	}
}

func TestFilterProperties_Empty(t *testing.T) {
	properties := createTestProperties()
	price := 9999.0
	filter := models.PropertyFilter{
		MinPrice: &price,
	}

	result := FilterProperties(properties, filter)
	if result==nil {
		t.Errorf("expected empty slice, got nil")
	}
	if len(result)!=0 {
		t.Errorf("expected 0 result got %d", len(result))
	}
}

func TestGetPropertyByID_Found(t *testing.T) {

	property := models.SourceProperty{
		ID: "TEST-ID",
		Images: []string{
			"a.jpg",
		},
	}

	sourceProperties = []models.SourceProperty{
		property,
	}

	result, err := GetPropertyByID("TEST-ID")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result == nil {
		t.Error("expected property but got nil")
	}

	if result.ID != "TEST-ID" {
		t.Errorf("expected TEST-ID got %s", result.ID)
	}
}

func TestGetPropertyByID_NotFound(t *testing.T) {

	sourceProperties = []models.SourceProperty{
		{
			ID: "TEST-ID",
		},
	}

	result, err := GetPropertyByID("UNKNOWN-ID")


	if err == nil {
		t.Error("expected error but got nil")
	}


	if result != nil {
		t.Error("expected nil property")
	}
}