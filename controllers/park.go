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

// Method : GET
//
// Return :
// {
//   "Result": 1,
//   "Err": "ok",
//   "Parks": [
//     {
//       "ParkName": "abcdd",
//       "ParkId": "4",
//       "Longitude": 100.01,
//       "Latitude": 99.91,
//       "Storey": 0,
//       "Empty": 0,
//       "Money": 5
//     },
//     {
//       "ParkName": "abcddded",
//       "ParkId": "5",
//       "Longitude": 100.01,
//       "Latitude": 99.91,
//       "Storey": 0,
//       "Empty": 0,
//       "Money": 5
//     }
//   ]
// }
// Result : 1. ok
//          0. error

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
