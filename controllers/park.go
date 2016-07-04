package controllers

import (
	"encoding/json"
	"fmt"
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
	var results []models.ParkResult
	num, e1 := models.FreeOutTimeSpace()
	if e1 == nil {
		fmt.Println(num)
		// TODO: recover the empty space +1
	}
	results, e := models.QueryParks(n)
	if e != nil {
		p.Data["json"] = &models.SearchResult{Result: 0, Err: e.Error()}
	} else {
		p.Data["json"] = &models.SearchResult{Result: 1, Err: "ok", Parks: results}
	}
	p.ServeJSON()
}

// @Title search by location
// @Description near by parkinglots
// @Success 200 {object} model.NearByResult
// @Failure 403 parkionglots not found
// @router /nearby [post]
func (p *ParkController) NearByParkingLots() {
	type Message struct {
		Longitude, Latitude float64
	}
	var m Message
	json.Unmarshal(p.Ctx.Input.RequestBody, &m)

	num, e1 := models.FreeOutTimeSpace()
	if e1 == nil {
		fmt.Println(num)
		// TODO: recover the empty space +1
	}

	var results []models.ParkResult
	results, e := models.NearByParks(m.Longitude, m.Latitude)
	if e != nil {
		p.Data["json"] = &models.NearByResult{Result: 0, Err: e.Error()}
	} else {
		p.Data["json"] = &models.NearByResult{Result: 1, Err: "ok", Parks: results}
	}
	// TODO: free the out of time parking space
	p.ServeJSON()
}
