package controllers

import (
	"encoding/json"

	"github.com/YoungsoonLee/backend_datainfra/libs"
	"github.com/YoungsoonLee/backend_datainfra/models"
)

type PaymentGatewayController struct {
	BaseController
}

// Post ...
func (p *PaymentGatewayController) Post() {

	var pg models.PaymentGateway
	err := json.Unmarshal(p.Ctx.Input.RequestBody, &pg)
	if err != nil {
		p.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	// TODO: validation

	// save to db
	pgid, err := models.AddPaymentGateway(pg)
	if err != nil {
		p.ResponseError(libs.ErrDatabase, err)
	}

	//success
	p.ResponseSuccess("pgid", pgid)
}
