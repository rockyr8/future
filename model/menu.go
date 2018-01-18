package model

import (
	"future/tool"
	db "future/database"
	//"fmt"
)

type Menu struct{
	ID string
	Nickname string
	Ico string
	Url string
}

//获取用户信息
func (m *Menu) GetList(id string) (rat string, err error) {
	var sql string
	if id!=""{
		sql= "SELECT * FROM go_menu_class where id=?"
	}else{
		sql = "SELECT * FROM go_menu_class where 1!=?"
	}
	rat, err = tool.DBResultTOJSON(sql,id)
	//fmt.Printf("sql=%s,err=%s\n",sql,err)
	return
}

//添加菜单
func (m *Menu) Add() (id int64, err error) {
	sql := "INSERT INTO go_menu_class (nickname,ico,url) VALUES (?,?,?)"
	rs, err := db.SqlDB.Exec(sql, m.Nickname,m.Ico,m.Url)
	// rs, err := db.SqlDB.Exec("call test1(?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

//修改菜单
func (m *Menu) Modify() (rows int64, err error) {
	sql := "UPDATE go_menu_class SET nickname=?,ico=?,url=? WHERE id=?"
	rs, err := db.SqlDB.Exec(sql, m.Nickname,m.Ico,m.Url,m.ID)
	// rs, err := db.SqlDB.Exec("call test1(?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	rows, err = rs.RowsAffected()
	return
}