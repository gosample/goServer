package controllers

import (
	"parkinglots/models"

	"github.com/astaxie/beego"
)

// Operations about Parks
type ParkController struct {
	beego.Controller
}

// @Title search by name
// @Description search parkinglots by parkName
// @Success 200 {object} model.SearchResult
// @Failure 403 parkionglots not found
// @router /search [get]
func (p *ParkController) SearchParkingLots() {
	var n string
	n = p.GetString("ParkName")
	var results []models.Result
	results, e := models.QueryParks(n)
	if e != nil {
		p.Data["json"] = &models.SearchResult{Result: 0, Err: e.Error()}
	} else {
		p.Data["json"] = &models.SearchResult{Result: 1, Err: "ok", Parks: results}
	}
	p.ServeJSON()
}
