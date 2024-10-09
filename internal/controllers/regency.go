package controllers

import (
	"mediversepreip/internal/models"

	"github.com/sev-2/raiden"
)

type RegencyControllers struct {
	raiden.ControllerBase
	Http  string `path:"/regency" type:"rest"`
	Model models.Regency
}
