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
	ChildNum string
}

type ChildMenu struct{
	ID string
	ClassID string //主菜单ID
	NickName string
	Authurl string
	Url string
	Sort string
	Valid string
}

//获取主菜单list信息
func (m *Menu) GetList() (rat string, err error) {

	sql := "SELECT * FROM go_menu_class where id>0"
	if m.ID !="" {
		sql += " and id=?"
		rat, err = tool.DBResultTOJSON(sql,m.ID)
		return
	}
	if m.ChildNum !="" {
		sql = "SELECT *,(SELECT COUNT(0) FROM go_menu m WHERE m.classID=a.id) AS num FROM go_menu_class a"
		rat, err = tool.DBResultTOJSON(sql)
		return
	}

	rat, err = tool.DBResultTOJSON(sql)
	return

}

//添加主菜单
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

//修改主菜单
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

//获取子list信息
func (m *ChildMenu) GetList() (rat string, err error) {
	sql := "SELECT IFNULL(b.nickname,'-') AS className,a.* FROM go_menu a LEFT JOIN go_menu_class b ON a.classID=b.id where 1=1"
	var args []interface{}
	if m.ID != "" {
		sql += " and a.id=?"
		args = append(args,m.ID)
	}
	if m.ClassID != "" {
		sql += " and a.classID=?"
		args = append(args,m.ClassID)
	}
	sql += " ORDER BY a.id DESC"
	rat, err = tool.DBResultTOJSON(sql,args...)
	//fmt.Printf("sql=%s,err=%s\n",sql,err)
	return
}

//添加子菜单
func (m *ChildMenu) Add() (id int64, err error) {
	sql := "INSERT INTO go_menu (classID,nickname,authurl,url,sort,valid) VALUES (?,?,?,?,?,?)"
	rs, err := db.SqlDB.Exec(sql, m.ClassID,m.NickName,m.Authurl,m.Url,m.Sort,m.Valid)
	// rs, err := db.SqlDB.Exec("call test1(?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

//修改子菜单
func (m *ChildMenu) Modify() (rows int64, err error) {
	sql := "UPDATE go_menu SET nickname=?,authurl=?,url=?,sort=?,valid=?,classID=? WHERE id=?"
	rs, err := db.SqlDB.Exec(sql, m.NickName,m.Authurl,m.Url,m.Sort,m.Valid,m.ClassID,m.ID)
	// rs, err := db.SqlDB.Exec("call test1(?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	rows, err = rs.RowsAffected()
	return
}