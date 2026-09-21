package controllers

import (
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
	c.Data["json"] = "Get all properties"
	c.ServeJSON()
}