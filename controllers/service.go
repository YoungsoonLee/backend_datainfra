package controllers

import (
	"encoding/json"

	"github.com/YoungsoonLee/backend_datainfra/libs"
	"github.com/YoungsoonLee/backend_datainfra/models"
)

type ServiceController struct {
	BaseController
}

// Post ...
func (s *ServiceController) Post() {

	var service models.Service
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &service)
	if err != nil {
		s.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	// TODO: validation

	// save to db
	sid, err := models.AddService(service)
	if err != nil {
		s.ResponseError(libs.ErrDatabase, err)
	}

	//success
	s.ResponseSuccess("sid", sid)
}
