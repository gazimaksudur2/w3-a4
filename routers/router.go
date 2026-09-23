// @APIVersion 1.0.0
// @Title Rental Property API
// @Description REST API for rental property search and lookup
// @BasePath /v1

package routers

import (
	"w3-a4/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSInclude(
			&controllers.PropertyController{},
		),
	)
	beego.AddNamespace(ns)
}