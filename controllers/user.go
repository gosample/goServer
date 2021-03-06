package controllers

import (
	"encoding/json"
	"parkinglots/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.RegisterResult
// @Failure 403 body is empty
// @router /register [post]
func (u *UserController) Register() {
	// TODO: replace `type struct` with model.User in the future.
	type Message struct {
		UserAccount, UserPwd, CarLicense string
	}
	var m Message
	json.Unmarshal(u.Ctx.Input.RequestBody, &m)

	e := models.AddUser(models.User{UserAcount: m.UserAccount, UserPwd: m.UserPwd})
	//TODO: check repeat CarLicense
	e1 := models.AddCar(models.Carofuser{UserAccount: m.UserAccount, CarLicense: m.CarLicense})
	if e != nil {
		u.Data["json"] = &models.RegisterResult{Result: 1, Err: e.Error() + e1.Error()}
	} else {
		u.Data["json"] = &models.RegisterResult{Result: 0, Err: "ok"}
	}
	u.ServeJSON()
}

// @Title login
// @Description Logs user into the system
// @Success 200 {object} model.LoginResult
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	// TODO: replace `type struct` with model.User in the future.
	type Message struct {
		UserAccount, UserPwd, CarLicense string
	}
	var m Message
	json.Unmarshal(u.Ctx.Input.RequestBody, &m)
	var results []models.CarResult
	results, e := models.QueryCar(m.UserAccount)

	token, e := models.UpdateUserToken(models.User{UserAcount: m.UserAccount, UserPwd: m.UserPwd})
	if e != nil {
		u.Data["json"] = &models.LoginResult{Result: 1, Err: e.Error()}
	} else {
		u.Data["json"] = &models.LoginResult{Result: 0, Err: "ok", Token: token, CarLicenses: results}
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

// @Title Hello
// @Success 200 {string} hello
// @router /:uid/hello [get]
func (u *UserController) Hello() {
	uid := u.GetString(":uid")
	u.Data["json"] = "Hello from " + uid
	u.ServeJSON()
}

// @Title AddCarLicense
// @Description user to add more carlicense
// @Success 200 {object} model.AddCarLicenseResult
// @Failure 403 add carlicense fail
// @router /addcarlicense [post]
func (u *UserController) AddCarLicense() {
	type Message struct {
		UserAccount, CarLicense string
	}
	var m Message
	json.Unmarshal(u.Ctx.Input.RequestBody, &m)
	e := models.AddCar(models.Carofuser{UserAccount: m.UserAccount, CarLicense: m.CarLicense})
	if e != nil {
		u.Data["json"] = &models.AddCarLicenseResult{Result: 1, Err: e.Error()}
	} else {
		u.Data["json"] = &models.AddCarLicenseResult{Result: 0, Err: "ok"}
	}
	u.ServeJSON()
}
