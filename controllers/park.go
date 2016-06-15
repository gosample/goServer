package controllers

import (
	"encoding/json"
	"parkinglots/models"

	"github.com/astaxie/beego"
)

// Operations about Parks
type ParkController struct {
	beego.Controller
}

// Method : GET
//
// Return :
// {
//     "Result": 1,
//     "Err": "ok",
//     "Parks": [
//         {
//             "ParkName": "鄞州体育馆停车场",
//             "ParkId": "3",
//             "Longitude": 121.561854,
//             "Latitude": 29.810474,
//             "StoreyNum": 100,
//             "EmptyNum": 50,
//             "Money": 5
//         },
//         {
//             "ParkName": "鄞州公园停车场",
//             "ParkId": "6",
//             "Longitude": 121.546649,
//             "Latitude": 29.811662,
//             "StoreyNum": 100,
//             "EmptyNum": 50,
//             "Money": 5
//         },
//         {
//             "ParkName": "鄞州区政府停车场",
//             "ParkId": "8",
//             "Longitude": 121.546616,
//             "Latitude": 29.817279,
//             "StoreyNum": 100,
//             "EmptyNum": 50,
//             "Money": 5
//         }
//     ]
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
// Body:
// {
//     "Longitude": 121.571636,
//     "Latitude": 29.816461
// }
// Return :
// {
//     "Result": 1,
//     "Err": "ok",
//     "Parks": [
//         {
//             "ParkName": "鄞州体育馆停车场",
//             "ParkId": "3",
//             "Longitude": 121.561854,
//             "Latitude": 29.810474,
//             "StoreyNum": 100,
//             "EmptyNum": 50,
//             "Money": 5
//         },
//         {
//             "ParkName": "鄞州公园停车场",
//             "ParkId": "6",
//             "Longitude": 121.546649,
//             "Latitude": 29.811662,
//             "StoreyNum": 100,
//             "EmptyNum": 50,
//             "Money": 5
//         },
//         {
//             "ParkName": "鄞州区政府停车场",
//             "ParkId": "8",
//             "Longitude": 121.546616,
//             "Latitude": 29.817279,
//             "StoreyNum": 100,
//             "EmptyNum": 50,
//             "Money": 5
//         }
//     ]
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

	var results []models.Result
	results, e := models.NearByParks(m.Longitude, m.Latitude)
	if e != nil {
		p.Data["json"] = &models.NearByResult{Result: 0, Err: e.Error()}
	} else {
		p.Data["json"] = &models.NearByResult{Result: 1, Err: "ok", Parks: results}
	}
	p.ServeJSON()
}
