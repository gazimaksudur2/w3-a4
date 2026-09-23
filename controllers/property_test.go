package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)


func createTestController(
	method string,
	path string,
) (*PropertyController, *httptest.ResponseRecorder) {


	req := httptest.NewRequest(
		method,
		path,
		nil,
	)

	w := httptest.NewRecorder()


	ctx := context.NewContext()

	ctx.Reset(
		w,
		req,
	)


	controller := &PropertyController{}

	controller.Controller.Init(
		ctx,
		"PropertyController",
		"GetAll",
		nil,
	)


	return controller,w
}



func TestPropertyController_GetAll_Success(t *testing.T){

	controller,w := createTestController(
		"GET",
		"/v1/properties",
	)


	controller.GetAll()


	if w.Code != 200 {
		t.Fatalf(
			"expected 200, got %d",
			w.Code,
		)
	}


	var response map[string]interface{}

	err := json.Unmarshal(
		w.Body.Bytes(),
		&response,
	)

	if err != nil {
		t.Fatal(err)
	}

}


func TestPropertyController_GetAll_InvalidPrice(t *testing.T){

	controller,w := createTestController(
		"GET",
		"/v1/properties?min_price=abc",
	)


	controller.GetAll()


	if w.Code != 400 {

		t.Fatalf(
			"expected 400 got %d",
			w.Code,
		)

	}

}


func TestPropertyController_GetAll_InvalidFeed(t *testing.T) {

	req := httptest.NewRequest(
		"GET",
		"/v1/properties?feed=99",
		nil,
	)

	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)


	controller := &PropertyController{
		Controller: beego.Controller{
			Ctx: ctx,
		},
	}


	controller.GetAll()


	if w.Code != 400 {
		t.Errorf(
			"expected 400, got %d",
			w.Code,
		)
	}

}

func TestPropertyController_GetByID_NotFound(t *testing.T){

	req := httptest.NewRequest(
		"GET",
		"/v1/properties/invalid-id",
		nil,
	)

	w := httptest.NewRecorder()


	ctx := context.NewContext()
	ctx.Reset(w, req)


	ctx.Input.SetParam(
		":id",
		"invalid-id",
	)


	controller := &PropertyController{
		Controller: beego.Controller{
			Ctx: ctx,
		},
	}


	controller.GetByID()


	if w.Code != 404 {
		t.Errorf(
			"expected 404 got %d",
			w.Code,
		)
	}

}