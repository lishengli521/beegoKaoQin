package models

import (
	"fmt"
	"mysqlgo/utils"
)

//获取侧边菜单
type Meun struct {
	Id   int     `db:"id"`
	Pid  int    `db:"pid"`
	Name  string `db:"name"`
	Url  string  `db:"url"`
	Icon  string  `db:"icon"`
	Flag  string  `db:"flag"`
	Sort  string  `db:"sort"`
	Children  []Meun
}

func GetMeunList(id int) []Meun {
   var meuns []Meun
	Db:=utils.GetDb()
	err :=Db.Select(&meuns,"select  id,pid,name,url,icon,flag,sort  FROM meun where pid=? ORDER BY sort ",id)
	if len(meuns)>0{
			DiGui(meuns)
	}
	if err != nil {
		fmt.Println("exec failed, ", err)
		return meuns
	}
	fmt.Println("select succ:", meuns)
	return meuns
}

func DiGui(meuns []Meun )  {
	Db:=utils.GetDb()
	for i:=0; i< len(meuns) ; i++{
		var children []Meun
		Db.Select(&children,"select  id,pid,name,url,icon,flag,sort  FROM meun where pid=? ORDER BY sort ",meuns[i].Id)
		meuns[i].Children=children
		if len(children)>0{
			DiGui(children)
		}
	}



}