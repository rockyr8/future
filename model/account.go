//用户管理的model，用于和数据库建立关系
package model

import (
	"fmt"
	"strconv"
	"time"

	. "future/logic"
	tool "future/tool"
	db "future/database"
)

//用户登录 成功则产生token 用于AuthMiddleWare和AuthLogicMiddleWare中间件验证 一个用户对应2个redis key value
func AccountLogin(uname,pwd string) (rat string,err error) {
	var uid,nickname,token string
	rows, err := db.SqlDB.Query("SELECT id,nickname FROM go_account WHERE uname=? AND pwd=? LIMIT 1",uname,pwd)
	defer rows.Close()
	if err != nil {
		return "",err
	}
	for rows.Next() {
		rows.Scan(&uid, &nickname)
	}
	if err = rows.Err(); err != nil {
		return "",err
	}
	if uid == "" {
		return "",err
	}

	//获取高级权限
	go GetAcceseAuth(uid,20,0)

	//生成原始 token
	token, err = tool.GenerateRandomString(20)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		return "",err
	}
	//拼接token 保证唯一性
	token += uid
	// fmt.Println(token)
	//加入到redis里面，用于基本验证。默认存储20分钟。
	if err := db.RedisSet(uid,token,60*20); err != nil{
		return "",err
	}

	rat = fmt.Sprintf(`[{"uid":"%s","token":"%s"}]`,uid,token)
	// fmt.Println(rat)
	return
}

//用户注销，删除redis数据，并向客户端发送指令，客户端收到指令删除对应的localStorage或者sessionStorage数据。
func AccountLoginOut(uid,ctoken string) error{
	// fmt.Printf("uid is %s, token is %s\n", uid, ctoken)
	if uid != "" && ctoken != "" {
		token := db.RedisGet(uid)
		if token == ctoken{
			err := db.RedisDel(uid)
			// err := db.RedisSet(uid,"",0)	
			go GetAcceseAuth(uid,0,1)
			return err
		}
	}

	return nil
	
}

//获取用户信息
func GetAccountList() (rat string,err error) {
	rat,err = tool.DBResultTOJSON("SELECT * FROM go_account")
	return 
}

func GetAccount1(uname,pwd string) (string,error) {
	slice,err := tool.DBResultDump("SELECT * FROM go_account")
	if err != nil{
		return "",err
	}
	// fmt.Println(slice)
	nick := slice[0]["nick"]
	unc := slice[0]["uname"].(string)+"~opoioe&http://wwcow.com.wzggw"
	valid,err := strconv.Atoi(slice[0]["valid"].(string))
	valid++

	ctime,err := strconv.Atoi(slice[0]["createtime"].(string))
	timeLayout := "2006-01-02 15:04:05" 
	dataTimeStr := time.Unix(int64(ctime), 0).Format(timeLayout)
	// nick = strings.Join([]string{nick, "~opoioe&http://wwcow.com.wzggw"}, "")
	
	
	for sl := range slice {
		for ssl,val := range slice[sl] {		
			// if wwe,ok := val.(string);ok{
			// 	wwe = val.(string) + "~opoioe&http://wwcow.com.wzggw"
			// 	fmt.Printf("%v\t%s\n", ssl,wwe)
			// 	continue
			// }
			// fmt.Printf("%v\t%s\n", ssl,val)

			switch val.(type){
				
				case string:
					fmt.Printf("字符串:: %v\t%s\n", ssl,val)
				
				case int:
					fmt.Printf("数字:: %v\t%s\n", ssl,val)
				
				default:
					fmt.Printf("%v\t%s\n", ssl,val)
				}
		}
		
	}

	fmt.Printf("%s\t%s\t%d\t%s\n", nick,unc,valid,dataTimeStr)

	return "",nil
}

//获取用户左边菜单
func GetMenu(uid string)(string,error){
	return tool.DBResultTOJSON(`SELECT DISTINCT a.nickname,url FROM go_menu a LEFT JOIN go_menu_role b ON a.id=b.menuID LEFT JOIN go_account_role c ON b.roleID=c.roleID 
		WHERE a.valid=0 AND c.accountID=? ORDER BY a.sort DESC`,uid)
}


