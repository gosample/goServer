package controllers

import (
	"fmt"
	"parkinglots/models"

	"github.com/astaxie/beego"
)

// Operations about ParkingSpace
type SpaceController struct {
	beego.Controller
}

// @Title get empty space
// @Description get empty space by parkId
// @Success 200 {object} model.SpaceResult
// @Failure 403 parkingspace not found
// @router /emptyspace [get]
func (s *SpaceController) GetEmptySpace() {
	var n string
	n = s.GetString("ParkId")
	fmt.Println(n)
	results, num, e := models.QueryEmptySpace(n)
	if e != nil {
		s.Data["json"] = &models.SpaceResult{Result: 1, Err: e.Error()}
	} else {
		s.Data["json"] = &models.SpaceResult{Result: 0, Num: num, Err: "ok", ParkSpaces: results}
	}
	s.ServeJSON()

}
