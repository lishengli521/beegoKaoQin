package models

import (
	"fmt"
	"mysqlgo/utils"
)

type User struct {
	Id   int     `db:"id"`
	Name string `db:"name"`
	Sex  string `db:"sex"`
	Pwd  string  `db:"pwd"`
}
func GetAll() []User {
	Db:=utils.GetDb()
	var user []User
	err := Db.Select(&user, "select id, name, sex, pwd from user")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return user
	}
	fmt.Println("select succ:", user)
	return user
}