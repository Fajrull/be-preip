package controllers

import (
	"mediversepreip/internal/models"

	"github.com/sev-2/raiden"
)

type ProvinceController struct {
	raiden.ControllerBase
	Http  string `path:"/province" type:"rest"`
	Model models.Province
}
