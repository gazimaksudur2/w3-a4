package services

import (
	"testing"
	"w3-a4/models"
)

func createTestProperties() []models.SourceProperty {
	return []models.SourceProperty{
		{
			ID:                   "1",
			Feed:                 11,
			Published:            false,
			USDPrice:             100,
			StarRating:           5,
			ReviewScoreGeneral:   4.5,
			NumberOfReview:       50,
			PropertyTypeCategory: "Hotel",
			BedroomCount:         3,
			AmenityCategories: []string{
				"Pool",
				"Internet",
			},
		},
		{
			ID:                   "2",
			Feed:                 12,
			Published:            true,
			USDPrice:             200,
			StarRating:           3,
			ReviewScoreGeneral:   3.5,
			NumberOfReview:       10,
			PropertyTypeCategory: "Apartment",
			BedroomCount:         1,
			AmenityCategories: []string{
				"Parking",
			},
		},
		{
			ID:                   "3",
			Feed:                 11,
			Published:            true,
			USDPrice:             300,
			StarRating:           4,
			ReviewScoreGeneral:   4.8,
			NumberOfReview:       100,
			PropertyTypeCategory: "Villa",
			BedroomCount:         4,
			AmenityCategories: []string{
				"Gym",
			},
		},
	}
}


func TestFilterProperties(t *testing.T) {

	properties := createTestProperties()

	tests := []struct {
		name          string
		filter        models.PropertyFilter
		expectedCount int
		expectedID    string
	}{

		{
			name: "filter by feed and published status",

			filter: func() models.PropertyFilter {

				published := false
				feed := 11

				return models.PropertyFilter{
					Feed:      &feed,
					Published: &published,
				}

			}(),

			expectedCount: 1,
			expectedID:    "1",
		},

		{
			name: "filter by price range",

			filter: func() models.PropertyFilter {

				min := 50.0
				max := 150.0

				return models.PropertyFilter{
					MinPrice: &min,
					MaxPrice: &max,
				}

			}(),

			expectedCount: 1,
		},

		{
			name: "filter by amenities using OR logic",

			filter: models.PropertyFilter{
				Amenities: []string{
					"Internet",
					"Parking",
				},
			},

			expectedCount: 2,
		},

		{
			name: "combined feed and amenities filter",

			filter: func() models.PropertyFilter {

				feed := 11

				return models.PropertyFilter{
					Feed: &feed,
					Amenities: []string{
						"Internet",
					},
				}

			}(),

			expectedCount: 1,
			expectedID:    "1",
		},

		{
			name: "no matching property",

			filter: func() models.PropertyFilter {

				min := 9999.0

				return models.PropertyFilter{
					MinPrice: &min,
				}

			}(),

			expectedCount: 0,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := FilterProperties(
				properties,
				tt.filter,
			)

			if len(result) != tt.expectedCount {

				t.Errorf(
					"expected %d results, got %d",
					tt.expectedCount,
					len(result),
				)
			}

			if tt.expectedID != "" &&
				result[0].ID != tt.expectedID {

				t.Errorf(
					"expected ID %s got %s",
					tt.expectedID,
					result[0].ID,
				)
			}

		})

	}

}