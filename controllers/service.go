package controllers

import (
	"parkinglots/models"

	"github.com/astaxie/beego"
)

// Operations about Service
type ServiceController struct {
	beego.Controller
}

// @Title get service
// @Description get service by user_account
// @Success 200 {object} model.ServiceResult
// @Failure 403 service not found
// @router /history [get]
func (s *ServiceController) GetHistoryService() {
	var n string
	n = s.GetString("UserAccount")
	var results []models.Service
	results, e := models.QueryService(n)
	if e != nil {
		s.Data["json"] = &models.ServiceResult{Result: 1, Err: e.Error()}
	} else {
		s.Data["json"] = &models.ServiceResult{Result: 0, Err: "ok", Services: results}
	}
	s.ServeJSON()
}
