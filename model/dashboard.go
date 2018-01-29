//统计
package model

import (
	db "future/database"
	"future/tool"
	"time"
	//"fmt"
)

type Dashboard struct{
	NewPlayer string `json:"newplayer"`
	LoginPlayer string `json:"loginplayer"`
	LoginCount string `json:"logincount"`
	OnlinePlayer string `json:"onlineplayer"`
	SginPlayer string `json:"signplayer"`
	CorePlayer string `json:"coreplayer"`
	OnlineTimeLen string `json:"onlinetimelen"`
	OnlineTimeLenAvg string `json:"onlinetimelenavg"`
	RechargeAmount string `json:"rechargeamount"`
	RechargeAmountTotal string `json:"totalrechargeamount"`
	Gold int64 `json:"gold"`
	Diamond int64 `json:"diamond"`
	GoldAndDiamond int64 `json:"golddiamond"`
	LowJackpot string `json:"ljackpot"`
	MidJackpot string `json:"mjackpot"`
	HighJackpot string `json:"hjackpot"`
}

func (d *Dashboard) GetList() (rat string) {
	timeLayout := "20060102"
	t := time.Now().Format(timeLayout)
	d.NewPlayer = getNum("SELECT COUNT(0) FROM kbe_accountinfos WHERE FROM_UNIXTIME(regtime,'%Y%m%d')='"+t+"'")
	d.LoginPlayer = getNum("SELECT COUNT(DISTINCT sm_avatarID) FROM tbl_SysLogs	WHERE sm_logType=10 AND FROM_UNIXTIME(sm_date,'%Y%m%d')='"+t+"'")
	d.LoginCount = getNum("SELECT COUNT(sm_avatarID) FROM tbl_SysLogs WHERE sm_logType=10 AND FROM_UNIXTIME(sm_date,'%Y%m%d')='"+t+"'")
	d.OnlinePlayer = getNum("SELECT count(0) FROM tbl_SysLogs WHERE sm_logType=89 AND FROM_UNIXTIME(sm_date,'%Y%m%d')='"+t+"'")
	d.SginPlayer = getNum("SELECT COUNT(DISTINCT sm_avatarid) FROM tbl_Sign WHERE FROM_UNIXTIME(sm_date,'%Y%m%d')='"+t+"'")
	d.CorePlayer = getNum(`SELECT FROM_UNIXTIME(c.regtime),FROM_UNIXTIME(c.lasttime),c.numlogin,a.id,a.sm_gold,a.sm_diamond,a.sm_name,c.bindata,c.accountName FROM tbl_Avatar a LEFT JOIN tbl_Account_characters_values b ON a.sm_dbid=b.sm_dbid LEFT JOIN kbe_accountinfos c ON b.parentID=c.entityDBID
								WHERE c.regtime>1466713420 AND a.id NOT IN (112,292,333,324,254,253)
								AND c.lasttime BETWEEN (UNIX_TIMESTAMP(NOW())-604800)  AND UNIX_TIMESTAMP(NOW())
								AND c.numlogin>7
								ORDER BY c.lasttime DESC`)
	d.getOnlineTimeLen()
	d.RechargeAmount = getNum("SELECT IFNULL(SUM(sm_price),0) FROM tbl_OnlinePay WHERE FROM_UNIXTIME(sm_date,'%Y%m%d')='"+t+"'")
	d.RechargeAmountTotal = getNum("SELECT IFNULL(SUM(sm_price),0) FROM tbl_OnlinePay")
	d.getGoldAndDiamond()
	d.getJackpot()
	rat,_ = tool.StructToJSON(&d)
	//fmt.Println(rat)
	return
}

func getNum(sql string) (num string) {
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
func (d *Dashboard) getOnlineTimeLen() (err error) {
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
	return
}

//公司收入，支出
func (d *Dashboard) getJackpot() (err error) {
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
	return
}

//奖池
func (d *Dashboard) getGoldAndDiamond() (err error) {
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
	return
}

