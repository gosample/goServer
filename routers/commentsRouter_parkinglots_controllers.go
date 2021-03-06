package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["parkinglots/controllers:BookController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:BookController"],
		beego.ControllerComments{
			"BookParkingLots",
			`/book`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:ObjectController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:ObjectController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:ObjectController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:ObjectController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:ObjectController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:ParkController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:ParkController"],
		beego.ControllerComments{
			"SearchParkingLots",
			`/search`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:ParkController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:ParkController"],
		beego.ControllerComments{
			"NearByParkingLots",
			`/nearby`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:ParkController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:ParkController"],
		beego.ControllerComments{
			"GetParkingLotsInfo",
			`/info`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:ServiceController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:ServiceController"],
		beego.ControllerComments{
			"GetHistoryService",
			`/history`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:SpaceController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:SpaceController"],
		beego.ControllerComments{
			"GetEmptySpace",
			`/emptyspace`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:UserController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:UserController"],
		beego.ControllerComments{
			"Register",
			`/register`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:UserController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:UserController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:UserController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:UserController"],
		beego.ControllerComments{
			"Hello",
			`/:uid/hello`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["parkinglots/controllers:UserController"] = append(beego.GlobalControllerRouter["parkinglots/controllers:UserController"],
		beego.ControllerComments{
			"AddCarLicense",
			`/addcarlicense`,
			[]string{"post"},
			nil})

}
