package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func init() {
	o = orm.NewOrm()
}

type ParkSpace struct {
	SpaceId string
	ParkId  string
	State   int64
	Floor   int64
}

func QueryEmptySpace(parkId string) ([]ParkSpace, int64, error) {
	var spaceResults []ParkSpace
	fmt.Println(parkId)
	num, e := o.Raw("SELECT `space_id`,`park_id`,`state`,`floor` FROM `parking`.parkingspace WHERE `park_id` = ? and `state` = 0;", parkId).QueryRows(&spaceResults)
	fmt.Println(num)
	if e != nil {
		return nil, 0, e
	}
	return spaceResults, num, nil
}
