//角色管理
package model

import (
	"future/tool"
	db "future/database"
	//"fmt"
	"fmt"
	"strings"
)

type Role struct{
	ID string
	NickName string
	Valid string
	UpdateIDs string //修改的角色菜单
}

func (r *Role) GetList() (rat string, err error) {
	sql := "SELECT * FROM go_role where 1=1"
	var args []interface{}
	if r.ID != "" {
		sql += " and id=?"
		args = append(args,r.ID)
	}
	rat, err = tool.DBResultTOJSON(sql,args...)
	//fmt.Printf("sql=%s,err=%s,r.ID=%s \n",sql,err,r.ID)
	return
}

func (r *Role) Add() (id int64, err error) {
	sql := "INSERT INTO go_role (nickname,valid) VALUES (?,?)"
	rs, err := db.SqlDB.Exec(sql, r.NickName,r.Valid)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (r *Role) Modify() (rows int64, err error) {
	sql := "UPDATE go_role SET nickname=?,valid=? WHERE id=?"
	rs, err := db.SqlDB.Exec(sql, r.NickName,r.Valid,r.ID)
	if err != nil {
		return
	}
	rows, err = rs.RowsAffected()
	return
}

//得到主菜单明细和子菜单个数
func (r *Role) GetMainMenuList() (rat string, err error) {
	sql:= "SELECT *,(SELECT COUNT(0) FROM go_menu m WHERE m.classID=a.id) AS num FROM go_menu_class a"
	rat, err = tool.DBResultTOJSON(sql)
	//fmt.Printf("sql=%s,err=%s\n",sql,err)
	return
}

//得到角色菜单
func(r *Role) GetRoleMenuList() (rat string,err error) {
	sql:= "SELECT menuid FROM go_menu_role WHERE roleID=?"
	rat, err = tool.DBResultTOJSON(sql,r.ID)
	//fmt.Printf("sql=%s,err=%s\n",sql,err)
	return
}

//修改角色菜单
func (r *Role) ModifyRoleMenu() (err error) {
	sql := "DELETE FROM go_menu_role WHERE roleID=?"
	rs, err := db.SqlDB.Exec(sql,r.ID)
	if err != nil {
		return
	}

	_, err = rs.RowsAffected()
	if err != nil {
		return
	}

	//批量设置权限
	insert, err := db.SqlDB.Prepare("INSERT INTO go_menu_role (menuID, roleID) VALUES(?,?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	begin, err := db.SqlDB.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _,value := range strings.Split(r.UpdateIDs,",") {
		if value != "" {
			_, err = begin.Stmt(insert).Exec(value,r.ID)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	err = begin.Commit()
	if err != nil {
		fmt.Println(err)
		begin.Rollback()
		return
	}

	return
}