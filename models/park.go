package models

import "github.com/astaxie/beego/orm"

func init() {
	o = orm.NewOrm()
}

type Parks struct {
	ParkId        string
	ParkName      string
	ParkPwd       string
	StoreyNum     int
	EmptyNum      int
	Longitude     float64
	Latitude      float64
	ParkIp        string
	Money         float64
	UnitPrice     float64
	BookUnitPrice float64
}

type Result struct {
	ParkName  string
	ParkId    string
	Longitude float64
	Latitude  float64
	StoreyNum int
	EmptyNum  int
	Money     int
}

func QueryParks(parkName string) ([]Result, error) {
	var results []Result
	_, e := o.Raw("SELECT `park_name`,`park_id`,`longitude`,`latitude`,`storey_num`,`empty_num`,`money` FROM `parks` WHERE `park_name` LIKE ?", "%"+parkName+"%").QueryRows(&results)
	if e != nil {
		return nil, e
	}
	return results, nil
}

func NearByParks(longitude, latitude float64) ([]Result, error) {
	var results []Result
	_, e := o.Raw("SELECT `park_name`,`park_id`,`longitude`,`latitude`,`storey_num`,`empty_num`,`money` FROM `parking`.parks WHERE 0.02 >= ST_LENGTH(ST_LINESTRINGFROMWKB(LineString(`parking`.parks.park_pt, point(?,?))));", longitude, latitude).QueryRows(&results)
	if e != nil {
		return nil, e
	}
	return results, nil
}
