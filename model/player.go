//角色管理
package model

import (
	"future/tool"
	db "future/database"
)

type Player struct{
	ID string
	NickName string
	UserName string
	Gold string
	Diamond string
	HeadImg string
	Valid string
	RegTime string
}

func (p *Player) GetList() (rat string, err error) {
	sql := "SELECT id,sm_regtime,sm_account,sm_password,sm_gold,sm_diamond,sm_head,sm_status,sm_nickname FROM tbl_user_info where 1=1"
	var args []interface{}
	if p.ID != "" {
		sql += " and id=?"
		args = append(args,p.ID)
	}
	if p.NickName != "" {
		sql += " and sm_nickname=?"
		args = append(args,p.NickName)
	}
	if p.UserName != "" {
		sql += " and sm_account=?"
		args = append(args,p.UserName)
	}
	if p.Gold != "" {
		sql += " and sm_gold=?"
		args = append(args,p.Gold)
	}
	if p.Diamond != "" {
		sql += " and sm_diamond=?"
		args = append(args,p.Diamond)
	}
	if p.HeadImg != "" {
		sql += " and sm_head=?"
		args = append(args,p.HeadImg)
	}
	if p.Valid != "" {
		sql += " and sm_status=?"
		args = append(args,p.Valid)
	}
	sql += "  ORDER BY id DESC"
	rat, err = tool.DBResultTOJSON(sql,args...)
	//fmt.Printf("sql=%s,err=%s,r.ID=%s \n",sql,err,r.ID)
	return
}

func (p *Player) Disable() (rows int64, err error) {
	sql := "UPDATE tbl_user_info SET sm_status=? WHERE id=?"
	rs, err := db.SqlDB.Exec(sql,p.Valid,p.ID)
	if err != nil {
		return
	}
	rows, err = rs.RowsAffected()
	return
}