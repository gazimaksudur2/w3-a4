package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["w3-a4/controllers:PropertyController"] = append(beego.GlobalControllerRouter["w3-a4/controllers:PropertyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/v1/properties`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["w3-a4/controllers:PropertyController"] = append(beego.GlobalControllerRouter["w3-a4/controllers:PropertyController"],
        beego.ControllerComments{
            Method: "GetByID",
            Router: `/v1/properties/{id}`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
