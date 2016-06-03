package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func init() {
	o = orm.NewOrm()
}

type Parks struct {
	ParkId            string
	ParkName          string
	ParkPwd           string
	ParkNumofStorey   int
	NumofEmptyParking int
	Longitude         float64
	Latitude          float64
	ParkIp            string
	Money             float64
	UnitPrice         float64
	BookUnitPrice     float64
}

type Result struct {
	ParkName  string
	ParkId    string
	Longitude float64
	Latitude  float64
	Storey    int
	Empty     int
	Money     int
}

func QueryParks(parkName string) ([]Result, error) {
	var results []Result
	fmt.Println(parkName)
	num, e := o.Raw("SELECT `park_name`,`park_id`,`longitude`,`latitude`,`park_num_of_storey`,`num_of_empty_parking`,`money` FROM `parks` WHERE `park_name` LIKE ?", "%"+parkName+"%").QueryRows(&results)
	fmt.Println(results)
	fmt.Println(num, e)
	if e != nil {
		return nil, e
	}
	return results, nil
}
