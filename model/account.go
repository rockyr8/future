//用户管理的model，用于和数据库建立关系
package model

import (
	"fmt"
	"strconv"
	"time"

	. "future/middleware"
	"future/tool"
	db "future/database"
	"future/common"
)

//用户信息
type Account struct {
	Uid      string
	UserName string
	PassWD   string
	NickName string
	Phone    string
	Tel      string
	SuperID  int
	proportions float32
	RoleID   string //角色ID
	Valid    string //是否启用
}

//用户登录 成功则产生token 用于AuthMiddleWare和AuthLogicMiddleWare中间件验证 一个用户对应2个redis key value
func AccountLogin(uname, pwd string) (rat string, err error) {
	var uid, nickname, token string
	rows, err := db.SqlDB.Query("SELECT id,nickname FROM go_account WHERE uname=? AND pwd=? LIMIT 1", uname, pwd)
	defer rows.Close()
	if err != nil {
		return "", err
	}
	for rows.Next() {
		rows.Scan(&uid, &nickname)
	}
	if err = rows.Err(); err != nil {
		return "", err
	}
	if uid == "" {
		return "", err
	}

	//获取高级权限  确保所有权限都拿到才能进入
	GetAcceseAuth(uid, common.RedisStorageTime, 0)

	//生成原始 token
	token, err = tool.GenerateRandomString(20)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		return "", err
	}
	//拼接token 保证唯一性
	token += uid
 	fmt.Println(token)
	//加入到redis里面，用于基本验证。默认存储20分钟。
	if err := db.RedisSet(uid, token, common.RedisStorageTime); err != nil {
		return "", err
	}

	rat = fmt.Sprintf(`[{"uid":"%s","token":"%s"}]`, uid, token)
	// fmt.Println(rat)
	return
}

//用户注销，删除redis数据，并向客户端发送指令，客户端收到指令删除对应的localStorage或者sessionStorage数据。
func AccountLoginOut(uid, ctoken string) error {
	// fmt.Printf("uid is %s, token is %s\n", uid, ctoken)
	if uid != "" && ctoken != "" {
		token := db.RedisGet(uid)
		if token == ctoken {
			err := db.RedisDel(uid)
			err = db.RedisDel(uid + "menu")
			// err := db.RedisSet(uid,"",0)	
			go GetAcceseAuth(uid, 0, 1)
			return err
		}
	}

	return nil

}

//获取用户信息
func (a *Account) GetList() (rat string, err error) {
	rat, err = tool.DBResultTOJSON("SELECT a.*,b.nickname AS roleName FROM go_account a LEFT JOIN go_role b ON a.roleID=b.id")
	return
}

//添加用户
func (a *Account) Add() (id int64, err error) {
	sql := "INSERT INTO go_account (uname,pwd,nickname,phone,tel,roleID,valid,createtime) VALUES (?,?,?,?,?,?,?,UNIX_TIMESTAMP(NOW()))"
	rs, err := db.SqlDB.Exec(sql, a.UserName, a.PassWD, a.NickName, a.Phone, a.Tel, a.RoleID, a.Valid)
	// rs, err := db.SqlDB.Exec("call test1(?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

//修改用户
func (a *Account) Modify() (rows int64, err error) {
	sql := "UPDATE go_account SET nickname=?,phone=?,tel=?,roleID=?,valid=?,lastModifyTime=UNIX_TIMESTAMP(NOW()) WHERE id=?"
	rs, err := db.SqlDB.Exec(sql, a.NickName, a.Phone, a.Tel, a.RoleID, a.Valid, a.Uid)
	// rs, err := db.SqlDB.Exec("call test1(?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	rows, err = rs.RowsAffected()
	return
}

//得到一天用户的信息记录
func (a *Account) GetDetail() (rat string, err error) {
	rat, err = tool.DBResultTOJSON("SELECT uname,nickname,phone,tel,roleID FROM go_account WHERE id=?", a.Uid)
	return
}

//修改密码
func (a *Account) ModifyPwd(oldpwd string) (rat string){
	sql := "call modifypwd(?,?,?)"
	rows, err := db.SqlDB.Query(sql,a.Uid,oldpwd,a.PassWD)
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		rows.Scan(&rat)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return

}

//获取用户左边菜单
func (a *Account) GetMenu() (string, error) {

	if rat := db.RedisGet(a.Uid + "menu"); rat == "" {

		str, err := tool.DBResultTOJSON(`SELECT DISTINCT d.nickname AS fathername,d.url AS fatherurl,d.ico,a.nickname,a.url FROM go_menu a LEFT JOIN go_menu_role b ON a.id=b.menuID LEFT JOIN go_account c ON b.roleID=c.roleID LEFT JOIN go_menu_class d ON a.classID=d.id
WHERE a.valid=1 AND a.classID>0 AND c.id=? ORDER BY a.classID,a.sort DESC`, a.Uid)
		//用户菜单存入redis中，减少数据库IO
		if err != nil {
			return "", err
		}
		db.RedisSet(a.Uid+"menu", str, common.RedisStorageTime)
		return str, nil

	} else {
		return rat, nil
	}

}

var accounts []Account
var sqlinsert string
var startID int

//生成等级关联
func CreateChild() (err error) {

	sql := "DELETE FROM go_account_child"
	rs, err := db.SqlDB.Exec(sql)
	if err != nil {
		return
	}

	_, err = rs.RowsAffected()
	if err != nil {
		return
	}


	sqlinsert = ""
	accounts = make([]Account, 0)
	rows, err := db.SqlDB.Query("SELECT id,superID FROM go_account")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var acc Account
		rows.Scan(&acc.Uid, &acc.SuperID)
		accounts = append(accounts, acc)
	}
	if err = rows.Err(); err != nil {
		return
	}

	//fmt.Println(accounts)

	for _,account := range accounts {
		startID ,_ = strconv.Atoi(account.Uid)
		findChildrens(startID)
	}
	if len(sqlinsert)>0 {
		sqlinsert = sqlinsert[0:len(sqlinsert)-1]
		//fmt.Println(sqlinsert)
		_, err := db.SqlDB.Exec("INSERT INTO go_account_child(accountID,childID) VALUES "+sqlinsert)
		return err
	}
	return

}

//查找自己的后代
func findChildrens(findID int){
	for _,account := range accounts {
		if findID == account.SuperID {
			//findID changeed findID=account.ID
			//fmt.Printf("%d,",account.Uid)
			id,_:= strconv.Atoi(account.Uid)
			sqlinsert += fmt.Sprintf("(%d,%d),",startID,id)
			findChildrens(id)
		}else {
			continue
		}
	}
}


//test
func GetAccount1(uname, pwd string) (string, error) {
	slice, err := tool.DBResultDump("SELECT * FROM go_account")
	if err != nil {
		return "", err
	}
	// fmt.Println(slice)
	nick := slice[0]["nick"]
	unc := slice[0]["uname"].(string) + "~opoioe&http://wwcow.com.wzggw"
	valid, err := strconv.Atoi(slice[0]["valid"].(string))
	valid++

	ctime, err := strconv.Atoi(slice[0]["createtime"].(string))
	timeLayout := "2006-01-02 15:04:05"
	dataTimeStr := time.Unix(int64(ctime), 0).Format(timeLayout)
	// nick = strings.Join([]string{nick, "~opoioe&http://wwcow.com.wzggw"}, "")

	for sl := range slice {
		for ssl, val := range slice[sl] {
			// if wwe,ok := val.(string);ok{
			// 	wwe = val.(string) + "~opoioe&http://wwcow.com.wzggw"
			// 	fmt.Printf("%v\t%s\n", ssl,wwe)
			// 	continue
			// }
			// fmt.Printf("%v\t%s\n", ssl,val)

			switch val.(type) {

			case string:
				fmt.Printf("字符串:: %v\t%s\n", ssl, val)

			case int:
				fmt.Printf("数字:: %v\t%s\n", ssl, val)

			default:
				fmt.Printf("%v\t%s\n", ssl, val)
			}
		}

	}

	fmt.Printf("%s\t%s\t%d\t%s\n", nick, unc, valid, dataTimeStr)

	return "", nil
}
