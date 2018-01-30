//统计
package model

import (
	db "future/database"
	"future/tool"
	"time"
	//"fmt"
)

type Dashboard struct{
	NewPlayer string `json:"newplayer,omitempty"`
	LoginPlayer string `json:"loginplayer,omitempty"`
	LoginCount string `json:"logincount,omitempty"`
	OnlinePlayer string `json:"onlineplayer,omitempty"`
	SginPlayer string `json:"signplayer,omitempty"`
	CorePlayer string `json:"coreplayer,omitempty"`
	OnlineTimeLen string `json:"onlinetimelen,omitempty"`
	OnlineTimeLenAvg string `json:"onlinetimelenavg,omitempty"`
	RechargeAmount string `json:"rechargeamount,omitempty"`
	RechargeAmountTotal string `json:"totalrechargeamount,omitempty"`
	Gold int64 `json:"gold,omitempty"`
	Diamond int64 `json:"diamond,omitempty"`
	GoldAndDiamond int64 `json:"golddiamond,omitempty"`
	LowJackpot string `json:"ljackpot,omitempty"`
	MidJackpot string `json:"mjackpot,omitempty"`
	HighJackpot string `json:"hjackpot,omitempty"`
}

func (d *Dashboard) GetList() (rat string) {
	timeLayout := "20060102"
	t := time.Now().Format(timeLayout)
	d.NewPlayer = GetNum("SELECT COUNT(0) FROM kbe_accountinfos WHERE FROM_UNIXTIME(regtime,'%Y%m%d')='"+t+"'")
	d.LoginPlayer = GetNum("SELECT COUNT(DISTINCT sm_avatarID) FROM tbl_SysLogs	WHERE sm_logType=10 AND FROM_UNIXTIME(sm_date,'%Y%m%d')='"+t+"'")
	d.LoginCount = GetNum("SELECT COUNT(sm_avatarID) FROM tbl_SysLogs WHERE sm_logType=10 AND FROM_UNIXTIME(sm_date,'%Y%m%d')='"+t+"'")
	d.OnlinePlayer = GetNum("SELECT count(0) FROM tbl_SysLogs WHERE sm_logType=89 AND FROM_UNIXTIME(sm_date,'%Y%m%d')='"+t+"'")
	d.SginPlayer = GetNum("SELECT COUNT(DISTINCT sm_avatarid) FROM tbl_Sign WHERE FROM_UNIXTIME(sm_date,'%Y%m%d')='"+t+"'")
	d.CorePlayer = GetNum(`SELECT count(0) FROM tbl_Avatar a LEFT JOIN tbl_Account_characters_values b ON a.sm_dbid=b.sm_dbid LEFT JOIN kbe_accountinfos c ON b.parentID=c.entityDBID
								WHERE c.regtime>1466713420 AND a.id NOT IN (112,292,333,324,254,253)
								AND c.lasttime BETWEEN (UNIX_TIMESTAMP(NOW())-604800)  AND UNIX_TIMESTAMP(NOW())
								AND c.numlogin>7
								ORDER BY c.lasttime DESC`)
	d.GetOnlineTimeLen()
	d.RechargeAmount = GetNum("SELECT IFNULL(SUM(sm_price),0) FROM tbl_OnlinePay WHERE FROM_UNIXTIME(sm_date,'%Y%m%d')='"+t+"'")
	d.RechargeAmountTotal = GetNum("SELECT IFNULL(SUM(sm_price),0) FROM tbl_OnlinePay")
	d.GetGoldAndDiamond()
	d.GetJackpot()
	rat,_ = tool.StructToJSON(&d)
	//fmt.Println(rat)
	return
}

func GetNum(sql string) (num string) {
	rows, err := db.SqlDB.Query(sql)
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		rows.Scan(&num)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

//玩家时长
func (d *Dashboard) GetOnlineTimeLen() (rat string,err error) {
	timeLayout := "20060102"
	t := time.Now().Format(timeLayout)
	rows, err := db.SqlDB.Query(`SELECT SUM(sm_logContent),SUM(sm_logContent)/COUNT(DISTINCT sm_avatarID) AS PJ FROM tbl_SysLogs
	WHERE sm_logType=10 AND FROM_UNIXTIME(sm_date,'%Y%m%d')=?`,t)
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		rows.Scan(&d.OnlineTimeLen,&d.OnlineTimeLenAvg)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return tool.StructToJSON(&d)
}

//奖池
func (d *Dashboard) GetJackpot() (rat string,err error) {
	var gold,diamond string
	var tp int
	rows, err := db.SqlDB.Query("SELECT sm_amount,sm_amountDam,sm_type FROM tbl_PoolAmount")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		rows.Scan(&gold,&diamond,&tp)
		if tp == 0 {
			d.LowJackpot = "金幣："+gold+",鑽石："+diamond
		}else if tp == 1{
			d.MidJackpot = "金幣："+gold+",鑽石："+diamond
		}else{
			d.HighJackpot = "金幣："+gold+",鑽石："+diamond
		}
	}
	if err = rows.Err(); err != nil {
		return
	}
	d.GoldAndDiamond = d.Gold+d.Diamond
	return tool.StructToJSON(&d)
}

//公司收入，支出
func (d *Dashboard) GetGoldAndDiamond() (rat string,err error) {
	timeLayout := "20060102"
	t := time.Now().Format(timeLayout)
	rows, err := db.SqlDB.Query("SELECT SUM(sm_coin_gold),SUM(sm_coin_diamonds*1000) FROM tbl_Payment WHERE FROM_UNIXTIME(sm_date,'%Y%m%d')=?",t)
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		rows.Scan(&d.Gold,&d.Diamond)
	}
	if err = rows.Err(); err != nil {
		return
	}
	d.GoldAndDiamond = d.Gold+d.Diamond
	return tool.StructToJSON(&d)
}