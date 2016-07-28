package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	o = orm.NewOrm()
}

type Bookservices struct {
	UserAccount string
	CarLicense  string
	Hours       float64
	ParkId      string
	SpaceId     string
}

func AddBookServices(b Bookservices) error {
	_, e := o.Raw("INSERT INTO bookservices (user_account, car_license, hours, park_id, space_id, book_time) VALUES (?, ?, ?, ?, ?, ?);", b.UserAccount, b.CarLicense, b.Hours, b.ParkId, b.SpaceId, time.Now()).Exec()
	if e != nil {
		return e
	}
	return nil
}

// func FreeOutTimeSpace() (int64, error) {
// 	res, e := o.Raw("DELETE FROM bookservices WHERE UNIX_TIMESTAMP(book_time) <= UNIX_TIMESTAMP()- hours*3600").Exec()
// 	if e != nil {
// 		return 0, e
// 	} else {
// 		num, _ := res.RowsAffected()
// 		return num, nil
// 	}
// }
