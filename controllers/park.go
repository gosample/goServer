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

// Method : GET
// Body:
// {
//     "Longitude": 121.571636,
//     "Latitude": 29.816461
// }
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

// Method : POST
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

	fmt.Println(m)

	var results []models.Result
	results, e := models.NearByParks(m.Longitude, m.Latitude)
	if e != nil {
		p.Data["json"] = &models.NearByResult{Result: 0, Err: e.Error()}
	} else {
		p.Data["json"] = &models.NearByResult{Result: 1, Err: "ok", Parks: results}
	}
	p.ServeJSON()
}
