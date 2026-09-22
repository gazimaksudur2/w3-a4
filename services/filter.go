package services

import "w3-a4/models"

func hasAmenity(propertyAmenities []string, required []string) bool {
	for _, wanted := range required {
		for _, available := range propertyAmenities {
			if wanted == available {
				return true
			}
		}
	}
	return false
}

func FilterProperties(
	properties []models.SourceProperty,
	filter models.PropertyFilter,
) []models.SourceProperty {
	result := make([]models.SourceProperty, 0)

	for _, property := range properties {
		if filter.MinPrice != nil && property.USDPrice < *filter.MinPrice {
			continue
		}
		if filter.MaxPrice!=nil && property.USDPrice > *filter.MaxPrice {
			continue
		}
		if filter.MinStarRating!=nil && property.StarRating < *filter.MinStarRating {
			continue
		}
		if filter.MinReviewScore!=nil && property.ReviewScoreGeneral < *filter.MinReviewScore {
			continue
		}
		if filter.MinReviews!=nil && property.NumberOfReview < *filter.MinReviews {
			continue
		}
		if filter.Published!=nil && property.Published != *filter.Published {
			continue
		}
		if filter.PropertyType!=nil && property.PropertyTypeCategory != *filter.PropertyType {
			continue
		}
		if filter.Feed!=nil && property.Feed != *filter.Feed {
			continue
		}
		if filter.MinBedroom!=nil && property.BedroomCount < *filter.MinBedroom {
			continue
		}
		if len(filter.Amenities)>0 {
			if !hasAmenity(property.AmenityCategories, filter.Amenities) {
				continue
			}
		}
		result = append(result, property)
	}
	return result
}