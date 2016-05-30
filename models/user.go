package models

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/astaxie/beego/orm"
)

var (
	o orm.Ormer
)

func init() {
	// orm.RegisterModel(new(User))
	o = orm.NewOrm()
}

type User struct {
	UserAcount string
	UserPwd    string
	PhoneNum   string
	Money      float64
	Token      string
}

func AddUser(u User) error {
	_, e := o.Raw("INSERT INTO users (user_account, user_pwd, phone_num) VALUES (?, ?, ?);", u.UserAcount, u.UserPwd, u.PhoneNum).Exec()
	if e != nil {
		return e
	}
	// TODO: add new model `carofuser`
	return nil
}

func UpdateUserToken(u User) (string, error) {
	var p string
	e := o.Raw("SELECT `user_pwd` FROM `users` WHERE `user_account` = ?;", u.UserAcount).QueryRow(&p)
	if e != nil {
		return "", e
	}
	if p != u.UserPwd {
		return "", fmt.Errorf("wrong password")
	}

	// TODO: add to token cache
	token := strconv.Itoa(int(rand.Int63())) + u.UserAcount
	_, e = o.Raw("UPDATE `users` SET `Token`= ? WHERE `user_account`= ?;", token, u.UserAcount).Exec()
	if e != nil {
		return "", e
	}
	return token, nil
}
