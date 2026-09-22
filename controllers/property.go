package controllers

import (
	"w3-a4/models"
	"w3-a4/services"

	beego "github.com/beego/beego/v2/server/web"
)


type PropertyController struct {
	beego.Controller
}

// @Title List all Properties
// @Description Get properties with filters
// @Success 200 {object} models.PropertyListResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @router /v1/properties [get]
func (c *PropertyController) GetAll() {
	properties, err := services.GetAllProperties()

	if err != nil {
		c.Data["json"] = map[string]string{
			"Error": err.Error(),
		}
		c.Ctx.ResponseWriter.WriteHeader(500)
		c.ServeJSON()
		return
	}

	c.Data["json"] = properties
	c.ServeJSON()
}

func(c *PropertyController) GetByID() {
	id := c.Ctx.Input.Param(":id")
	property, err := services.GetPropertyByID(id)

	if err!=nil {
		c.Data["json"] = models.ErrorResponse{
			Error: err.Error(),
		}
		c.Ctx.ResponseWriter.WriteHeader(500)
		c.ServeJSON()
		return
	}

	if property == nil {
		c.Data["json"] = models.ErrorResponse{
			Error: "Property not found for the provided ID",
		}
		c.Ctx.ResponseWriter.WriteHeader(404)
		c.ServeJSON()
		return
	}
	c.Data["json"] = property
	c.ServeJSON()
}