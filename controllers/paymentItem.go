package controllers

import (
	"encoding/json"

	"github.com/YoungsoonLee/backend_datainfra/libs"
	"github.com/YoungsoonLee/backend_datainfra/models"
)

type PaymentItemController struct {
	BaseController
}

// Post ...
func (p *PaymentItemController) Post() {

	var pi models.PaymentItem

	err := json.Unmarshal(p.Ctx.Input.RequestBody, &pi)
	if err != nil {
		p.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	// TODO: validation

	// save to db
	itemid, err := models.AddPaymentItem(pi)
	if err != nil {
		p.ResponseError(libs.ErrDatabase, err)
	}

	//success
	p.ResponseSuccess("ItemID", itemid)
}
