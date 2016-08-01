package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func init() {
	o = orm.NewOrm()
}

type Parks struct {
	ParkId        string
	ParkName      string
	ParkPwd       string
	StoreyNum     int64
	EmptyNum      int64
	Longitude     float64
	Latitude      float64
	ParkIp        string
	Money         float64
	UnitPrice     float64
	BookUnitPrice float64
	ParkImg       string
	ParkAddress   string
}

type ParkResult struct {
	ParkName  string
	ParkId    string
	Longitude float64
	Latitude  float64
	StoreyNum int64
	EmptyNum  int64
	Money     float64
	ParkImg   string
}

func QueryParks(parkName string) ([]ParkResult, error) {
	var parkResults []ParkResult
	_, e := o.Raw("SELECT `park_name`,`park_id`,`longitude`,`latitude`,`storey_num`,`empty_num`,`money`,`park_img` FROM `parks` WHERE `park_name` LIKE ?", "%"+parkName+"%").QueryRows(&parkResults)
	if e != nil {
		return nil, e
	}
	return parkResults, nil
}

func NearByParks(longitude, latitude float64) ([]ParkResult, error) {
	var parkResults []ParkResult
	_, e := o.Raw("SELECT `park_name`,`park_id`,`longitude`,`latitude`,`storey_num`,`empty_num`,`money`,`park_img` FROM `parking`.parks WHERE 0.02 >= ST_LENGTH(ST_LINESTRINGFROMWKB(LineString(`parking`.parks.park_pt, point(?,?))));", longitude, latitude).QueryRows(&parkResults)
	if e != nil {
		return nil, e
	}
	return parkResults, nil
}

func BookParkingLot(parkId string) error {

	fmt.Println(parkId)

	var emptyNum int
	e := o.Raw("SELECT `empty_num` FROM `parks` WHERE `park_id` = ?;", parkId).QueryRow(&emptyNum)
	if e != nil {
		return e
	}
	fmt.Println(emptyNum)

	if emptyNum <= 0 {
		return fmt.Errorf("no empty parking space")
	}
	return nil
}

func GetInfoById(parkId string) (ParkResult, error) {
	var parkResult ParkResult
	e := o.Raw("SELECT `park_name`,`park_id`,`longitude`,`latitude`,`storey_num`,`empty_num`,`money`,`park_img` FROM `parking`.parks WHERE `park_id` = ?;", parkId).QueryRow(&parkResult)
	if e != nil {
		return parkResult, e
	}
	return parkResult, nil
}
