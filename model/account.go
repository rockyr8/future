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
	"strings"
)

//用户信息
type Account struct {
	ID          int
	Uid         string
	UserName    string
	PassWD      string
	NickName    string
	Phone       string
	Tel         string
	SuperID     int
	Proportions float64 //分成比例
	RoleID      string  //角色ID
	Valid       string  //是否启用
	Logintime   string  //登录时间，范围
	Createtime  string  //注册时间，范围
}

//账户银行表
type AccountBank struct {
	Account
	AccountID      string
	AccountName    string
	AccountCardNum string
	OpenBank       string
	BranchBank     string
	Balance        float64
}

//销售表，分成直接来源
type SaleInfo struct {
	Account
	AccountID int
	SaleAmt   float64
}

//销售集合
var saleinfos []SaleInfo

//用户登录 成功则产生token 用于AuthMiddleWare和AuthLogicMiddleWare中间件验证 一个用户对应2个redis key value
func AccountLogin(uname, pwd string) (rat string, err error) {
	var uid, nickname, roleID, token string
	rows, err := db.SqlDB.Query("SELECT id,nickname,roleID FROM go_account WHERE uname=? AND pwd=? LIMIT 1", uname, pwd)
	defer rows.Close()
	if err != nil {
		return "", err
	}
	for rows.Next() {
		rows.Scan(&uid, &nickname, &roleID)
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
	//fmt.Println(token)
	//加入到redis里面，用于基本验证。默认存储20分钟。
	if err := db.RedisSet(uid, token, common.RedisStorageTime); err != nil {
		return "", err
	}
	//加入角色ID
	if err := db.RedisSet(uid+"role", roleID, common.RedisStorageTime); err != nil {
		return "", err
	}

	rat = fmt.Sprintf(`[{"uid":"%s","token":"%s"}]`, uid, token)
	// fmt.Println(rat)
	go updateLogintime(uid) //更新登录时间
	return
}

func updateLogintime(uid string) {
	sql := "UPDATE go_account SET lastlogintime=UNIX_TIMESTAMP(NOW()) WHERE id=?"
	_, err := db.SqlDB.Exec(sql, uid)
	if err != nil {
		return
	}
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
	var args []interface{}
	sql := "SELECT a.*,b.nickname AS roleName,IFNULL((SELECT nickname FROM go_account WHERE id=a.superID),'-') AS supername FROM go_account_child gc LEFT JOIN go_account a ON gc.childID=a.id LEFT JOIN go_role b ON a.roleID=b.id WHERE gc.accountID=? "
	args = append(args, a.Uid)
	if a.UserName != "" {
		sql += " and a.uname=?"
		args = append(args, a.UserName)
	}
	if a.NickName != "" {
		sql += " and a.nickname=?"
		args = append(args, a.NickName)
	}
	if a.Phone != "" {
		sql += " and a.phone=?"
		args = append(args, a.Phone)
	}
	if a.Tel != "" {
		sql += " and a.tel=?"
		args = append(args, a.Tel)
	}
	if a.Valid != "" {
		sql += " and a.Valid=?"
		args = append(args, a.Tel)
	}
	if a.RoleID != "" {
		sql += " and a.roleID=?"
		args = append(args, a.RoleID)
	}
	if a.Logintime != "" {
		sql += " and a.lastLoginTime BETWEEN ? AND ?"
		timerange := strings.Split(a.Logintime, "~")
		args = append(args, timerange[0], timerange[1])
	}
	if a.Createtime != "" {
		sql += " and a.createtime BETWEEN ? AND ?"
		timerange := strings.Split(a.Createtime, "~")
		args = append(args, timerange[0], timerange[1])
	}
	rat, err = tool.DBResultTOJSON(sql, args...)
	return
}

//添加用户
func (a *Account) Add() (id int64, err error) {
	sql := "INSERT INTO go_account (uname,pwd,nickname,phone,tel,roleID,valid,createtime,proportions,superID) VALUES (?,?,?,?,?,?,?,UNIX_TIMESTAMP(NOW()),?,?)"
	rs, err := db.SqlDB.Exec(sql, a.UserName, a.PassWD, a.NickName, a.Phone, a.Tel, a.RoleID, a.Valid, a.Proportions, a.ID)
	// rs, err := db.SqlDB.Exec("call test1(?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	if id > 0 {
		addBank(a.Uid) //添加银行信息
		CreateChild()  //更新子账户
	}
	return
}

//修改用户
func (a *Account) Modify() (rows int64, err error) {
	sql := "UPDATE go_account SET nickname=?,phone=?,tel=?,roleID=?,valid=?,lastModifyTime=UNIX_TIMESTAMP(NOW()),proportions=? WHERE id=?"
	rs, err := db.SqlDB.Exec(sql, a.NickName, a.Phone, a.Tel, a.RoleID, a.Valid, a.Proportions, a.Uid)
	// rs, err := db.SqlDB.Exec("call test1(?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	rows, err = rs.RowsAffected()
	return
}

//得到一个用户的信息记录
func (a *Account) GetDetail() (rat string, err error) {
	rat, err = tool.DBResultTOJSON("SELECT uname,nickname,phone,tel,roleID,proportions FROM go_account WHERE id=?", a.Uid)
	return
}

//得到一个用户银行信息记录
func (a *AccountBank) GetBankDetail() (rat string, err error) {
	rat, err = tool.DBResultTOJSON("SELECT * FROM go_account_bank WHERE accountID=?", a.AccountID)
	return
}

//添加银行信息
func addBank(uid string) (err error) {
	sql := "INSERT INTO go_account_bank (accountID) VALUES(?)"
	_, err = db.SqlDB.Exec(sql, uid)
	if err != nil {
		return
	}
	return
}

//修改银行信息
func (a *AccountBank) UpdateBank() (rows int64, err error) {
	sql := "UPDATE go_account_bank SET accountName=?,accountCardNum=?,openBank=?,branchBank=? WHERE accountID=?"
	rs, err := db.SqlDB.Exec(sql, a.AccountName, a.AccountCardNum, a.OpenBank, a.BranchBank, a.AccountID)
	if err != nil {
		return
	}
	rows, err = rs.RowsAffected()
	if err != nil {
		return
	}

	sql = "UPDATE go_account SET nickname=?,phone=?,tel=?,lastModifyTime=UNIX_TIMESTAMP(NOW()) WHERE id=?"
	rs, err = db.SqlDB.Exec(sql, a.Account.NickName, a.Account.Phone, a.Account.Tel, a.AccountID)
	if err != nil {
		return
	}
	rows, err = rs.RowsAffected()
	return
}

//修改密码
func (a *Account) ModifyPwd(oldpwd string) (rat string) {
	sql := "call modifypwd(?,?,?)"
	rows, err := db.SqlDB.Query(sql, a.Uid, oldpwd, a.PassWD)
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

//使用全局变量前，请重新赋零值。不赋值，会沿用上次使用的值。
var accounts []Account
var sqlinsert string
var startID int

//生成等级关联 下属账户
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
		rows.Scan(&acc.ID, &acc.SuperID)
		accounts = append(accounts, acc)
	}
	if err = rows.Err(); err != nil {
		return
	}

	//fmt.Println(accounts)

	for _, account := range accounts {
		//startID, _ = strconv.Atoi(account.Uid)
		startID = account.ID
		sqlinsert += fmt.Sprintf("(%d,%d),", startID, startID)
		findChildrens(startID)
	}
	if len(sqlinsert) > 0 {
		sqlinsert = sqlinsert[0:len(sqlinsert)-1]
		//fmt.Println(sqlinsert)
		_, err := db.SqlDB.Exec("INSERT INTO go_account_child(accountID,childID) VALUES " + sqlinsert)
		return err
	}
	return

}

//查找自己的后代
func findChildrens(superID int) {
	for _, account := range accounts {
		if superID == account.SuperID {
			//findID changeed findID=account.ID
			//fmt.Printf("%d,",account.Uid)
			//id, _ := strconv.Atoi(account.Uid)
			id := account.ID
			sqlinsert += fmt.Sprintf("(%d,%d),", startID, id)
			findChildrens(id)
		} else {
			continue
		}
	}
}

var yesTimeUnix int64
//分成表 生成昨天的分成金额
func CreateSettlement() (err error) {

	//设置查询的时间:当前是查昨天整天数据
	nTime := time.Now()
	yesTimeUnix = nTime.AddDate(0, 0, 0).Unix()
	yesTime := nTime.AddDate(0, 0, 0).Format("20060102")

	sql := "DELETE FROM go_account_settlement WHERE FROM_UNIXTIME(settledate,'%Y%m%d')=?"
	rs, err := db.SqlDB.Exec(sql, yesTime)
	if err != nil {
		return
	}

	_, err = rs.RowsAffected()
	if err != nil {
		return
	}

	/***********账户表***********/
	accounts = make([]Account, 0)
	rows, err := db.SqlDB.Query("SELECT id,superID,proportions FROM go_account")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var acc Account
		rows.Scan(&acc.ID, &acc.SuperID, &acc.Proportions)
		accounts = append(accounts, acc)
	}
	if err = rows.Err(); err != nil {
		return
	}

	/************销售表**********/
	saleinfos = make([]SaleInfo, 0)
	rows, err = db.SqlDB.Query("SELECT a.accountID,a.saleamt,b.superID,b.proportions FROM go_account_sale a LEFT JOIN go_account b ON a.accountID=b.id WHERE FROM_UNIXTIME(saledate,'%Y%m%d')=?", yesTime)
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var sale SaleInfo
		rows.Scan(&sale.AccountID, &sale.SaleAmt, &sale.Account.SuperID, &sale.Account.Proportions)
		saleinfos = append(saleinfos, sale)
	}
	if err = rows.Err(); err != nil {
		return
	}

	//fmt.Println(666, len(saleinfos))
	if len(saleinfos) > 0 {
		sqlinsert = ""
		for _, sale := range saleinfos {
			startID = sale.AccountID
			proportions := sale.Account.Proportions
			superID := sale.Account.SuperID
			selfmoney := sale.SaleAmt * proportions / 100.0
			supermoney := sale.SaleAmt * (100 - proportions) / 100.0
			sqlinsert += fmt.Sprintf("(%d,%.2f,%d,%d,0,%.2f),", sale.AccountID, selfmoney, time.Now().Unix(), superID, proportions)
			findSuperShare(superID, supermoney)
		}

		sqlinsert = sqlinsert[0:len(sqlinsert)-1]
		//fmt.Println(sqlinsert)
		_, err = db.SqlDB.Exec("INSERT INTO go_account_settlement (accountID,amt,settledate,superID,childID,proportions) VALUES " + sqlinsert)
		//fmt.Println(err)
	}
	return

}

//查找自己的祖先并且分账 不允许出现{ID:1,superID:3}{ID:3,superID:1} 这种情况,避免死循环
func findSuperShare(superID int, superMoney float64) {
	for _, account := range accounts {
		if superID == account.ID {
			minusmoney := superMoney - superMoney*account.Proportions/100.0
			selfmoney := superMoney * account.Proportions / 100.0
			//fmt.Printf("(%d,%d,%d,%.2f)\n", startID, account.ID, account.SuperID, account.Proportions)
			//fmt.Printf("账号：%d=%.2f\n\n",account.ID,selfmoney)
			sqlinsert += fmt.Sprintf("(%d,%.2f,%d,0,%d,%.2f),", account.ID, selfmoney, time.Now().Unix(), startID, account.Proportions)
			findSuperShare(account.SuperID, minusmoney)
		} else {
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
