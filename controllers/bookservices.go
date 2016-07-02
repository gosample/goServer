package controllers

import (
	"encoding/json"
	"parkinglots/models"

	"github.com/astaxie/beego"
)

// Operations about Bookservices
type BookController struct {
	beego.Controller
}

// @Title book parkinglots
// @Description book parkinglots
// @Success 200 {object} model.BookResult
// @Failure 403 parkionglots not empty
// @router /book [book]
func (b *BookController) BookParkingLots() {
	type Message struct {
		UserAccount, CarLicense, ParkId string
		Hours                           float64
	}
	var m Message
	json.Unmarshal(b.Ctx.Input.RequestBody, &m)

	e := models.AddBookServices(models.Bookservices{UserAccount: m.UserAccount, CarLicense: m.CarLicense, Hours: m.Hours, ParkId: m.ParkId})
	if e != nil {
		b.Data["json"] = &models.BookResult{Result: 1, Err: e.Error()}
	} else {
		b.Data["json"] = &models.BookResult{Result: 0, Err: "ok"}
	}
	//TODO:update empty_num in park table -1
	b.ServeJSON()
}
