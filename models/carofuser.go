package models

import "github.com/astaxie/beego/orm"

func init() {
	o = orm.NewOrm()
}

type Carofuser struct {
	UserAccount string
	CarLicense  string
}
type CarResult struct {
	CarLicense string
}

func AddCar(c Carofuser) error {
	_, e := o.Raw("INSERT INTO carofuser (user_account, car_license) VALUES (?, ?);", c.UserAccount, c.CarLicense).Exec()
	if e != nil {
		return e
	}
	return nil
}

func QueryCar(userAccount string) ([]CarResult, error) {
	var carResults []CarResult
	_, e := o.Raw("SELECT `car_license` FROM `parking`.carofuser WHERE `user_account` = ?", userAccount).QueryRows(&carResults)
	if e != nil {
		return nil, e
	}
	return carResults, nil
}
