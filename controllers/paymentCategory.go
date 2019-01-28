package controllers

import (
	"encoding/json"

	"github.com/YoungsoonLee/backend_datainfra/libs"
	"github.com/YoungsoonLee/backend_datainfra/models"
)

type PaymentCategoryController struct {
	BaseController
}

// Post ...
func (p *PaymentCategoryController) Post() {

	var pc models.PaymentCategory

	err := json.Unmarshal(p.Ctx.Input.RequestBody, &pc)
	if err != nil {
		p.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	// TODO: validation

	// save to db
	pcid, err := models.AddPaymentCategory(pc)
	if err != nil {
		p.ResponseError(libs.ErrDatabase, err)
	}

	//success
	p.ResponseSuccess("pcid", pcid)
}
