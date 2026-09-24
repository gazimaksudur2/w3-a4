package services

import (
	"testing"
	"w3-a4/models"
)

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