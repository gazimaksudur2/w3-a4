package controllers

import (
	"errors"
	"strconv"
	"strings"
	"w3-a4/models"
	"w3-a4/services"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	beego.Controller
}

// @Title List Properties
// @Description Get rental properties with optional filters
// @Param min_price query number false "Minimum price"
// @Param max_price query number false "Maximum price"
// @Param min_review_score query number false "Minimum review score"
// @Param feed query int false "Property feed (11,12,22,24)"
// @Param published query bool false "Published status"
// @Param property_type query string false "Property type"
// @Param min_star_rating query int false "Minimum star rating"
// @Param min_reviews query int false "Minimum number of reviews"
// @Param min_bedroom query int false "Minimum bedrooms"
// @Param amenities query string false "Comma separated amenities"
// @Param limit query int false "Maximum number of results"
// @Success 200 {object} models.PropertyListResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @router /v1/properties [get]
func (c *PropertyController) GetAll() {
	filter := models.PropertyFilter{}

	if value := c.GetString("min_price"); value != "" {
		price, err := strconv.ParseFloat(value, 64)
		if err != nil || price < 0 {
			c.Data["json"] = models.ErrorResponse{
				Error: "invalid min_price",
			}
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}
		filter.MinPrice = &price
	}

	if value := c.GetString("max_price"); value != "" {
		price, err := strconv.ParseFloat(value, 64)
		if err != nil || price < 0 {
			c.Data["json"] = models.ErrorResponse{
				Error: "invalid max_price",
			}
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}
		filter.MaxPrice = &price
	}

	if value := c.GetString("feed"); value != "" {
		feed, err := strconv.Atoi(value)
		if err != nil {
			c.Data["json"] = models.ErrorResponse{
				Error: "invalid feed",
			}
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}

		validFeeds := map[int]bool{
			11: true,
			12: true,
			22: true,
			24: true,
		}

		if !validFeeds[feed] {
			c.Data["json"] = models.ErrorResponse{
				Error: "invalid feed value",
			}
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}
		filter.Feed = &feed
	}

	if value := c.GetString("published"); value != "" {
		published, err := strconv.ParseBool(value)
		if err != nil {
			c.Data["json"] = models.ErrorResponse{
				Error: "invalid published value",
			}
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}
		filter.Published = &published
	}

	if value := c.GetString("property_type"); value != "" {
		validTypes := map[string]bool{
			"Hotel":     true,
			"House":     true,
			"Apartment": true,
			"Villa":     true,
			"Resort":    true,
			"Hostel":    true,
		}

		if !validTypes[value] {
			c.Data["json"] = models.ErrorResponse{
				Error: "invalid property_type",
			}
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}
		filter.PropertyType = &value
	}

	if value := c.GetString("min_star_rating"); value != "" {
		rating, err := strconv.Atoi(value)
		if err != nil || rating < 0 {
			c.Data["json"] = models.ErrorResponse{
				Error: "invalid min_star_rating",
			}
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}

		filter.MinStarRating = &rating
	}

	if value := c.GetString("min_review_score"); value != "" {
		score, err := strconv.ParseFloat(value, 64)
		if err != nil || score < 0 {
			c.Data["json"] = models.ErrorResponse{
				Error: "invalid min_review_score",
			}
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}
		filter.MinReviewScore = &score
	}

	if value := c.GetString("min_reviews"); value != "" {
		reviews, err := strconv.Atoi(value)
		if err != nil || reviews < 0 {
			c.Data["json"] = models.ErrorResponse{
				Error: "invalid min_reviews",
			}
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}
		filter.MinReviews = &reviews
	}

	if value := c.GetString("min_bedroom"); value != "" {
		bedroom, err := strconv.Atoi(value)
		if err != nil || bedroom < 0 {
			c.Data["json"] = models.ErrorResponse{
				Error: "invalid min_bedroom",
			}
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}
		filter.MinBedroom = &bedroom
	}

	if value := c.GetString("amenities"); value != "" {
		filter.Amenities = strings.Split(value, ",")
	}

	if value := c.GetString("limit"); value != "" {
		limit, err := strconv.Atoi(value)
		if err != nil || limit < 0 {
			c.Data["json"] = models.ErrorResponse{
				Error: "invalid limit",
			}
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.ServeJSON()
			return
		}
		filter.Limit = limit
	}

	response, err := services.ListProperties(filter)
	if err != nil {
		c.Data["json"] = models.ErrorResponse{
			Error: err.Error(),
		}
		c.Ctx.ResponseWriter.WriteHeader(500)
		c.ServeJSON()
		return
	}
	c.Data["json"] = response
	c.ServeJSON()
}

// @Title Get Property By ID
// @Description Get a single rental property by ID
// @Param id path string true "Property ID"
// @Success 200 {object} models.PropertyResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @router /v1/properties/{id} [get]
func (c *PropertyController) GetByID() {
	id := c.Ctx.Input.Param(":id")
	property, err := services.GetPropertyByID(id)

	if err != nil {
		status := 500
		if errors.Is(err, services.ErrPropertyNotFound) {
			status = 404
		}
		c.Data["json"] = models.ErrorResponse{
			Error: err.Error(),
		}
		c.Ctx.ResponseWriter.WriteHeader(status)
		c.ServeJSON()
		return
	}

	c.Data["json"] = property
	c.ServeJSON()
}
