package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	o = orm.NewOrm()
}

type Service struct {
	UserAccount string
	CarLicense  string
	Hours       float64
	ParkId      string
	StartTime   time.Time
	ExitTime    time.Time
	TotalMoney  float64
}

func QueryService(userAccount string) ([]Service, error) {
	var services []Service
	_, e := o.Raw("SELECT `user_account`,`car_license`,`hours`,`park_id`,`start_time`,`exit_time`,`total_money` FROM `service` WHERE `user_account` = ?", userAccount).QueryRows(&services)
	if e != nil {
		return nil, e
	}
	return services, nil
}
