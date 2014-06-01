package controllers

import (
	"github.com/ulricqin/gen-winning-numbers/models"
	"strings"
)

type ComputeController struct {
	AdminController
}

func (this *ComputeController) Model2() {
	result, err := models.ComputeModel2()
	if err != nil {
		this.Ctx.WriteString(err.Error())
	} else {
		this.Ctx.WriteString(strings.Join(result, " "))
	}
}
