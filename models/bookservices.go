package models

import "github.com/astaxie/beego/orm"

func init() {
	o = orm.NewOrm()
}

type Bookservices struct {
	UserAccount string
	CarLicense  string
	Hours       float64
	ParkId      string
}

func AddBookServices(b Bookservices) error {
	_, e := o.Raw("INSERT INTO bookservices (user_account, car_license, hours, park_id) VALUES (?, ?, ?, ?);", b.UserAccount, b.CarLicense, b.Hours, b.ParkId).Exec()
	if e != nil {
		return e
	}
	return nil
}
