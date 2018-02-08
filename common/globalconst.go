package common

import "time"

const (

	PermissionVerification string = "" //权限验证关闭=>字符不能空,启用=>空字符串

	RedisStorageTime time.Duration = 50*60 //默认redis存储时间x分钟,第一个数字是分钟
	RenewalTime time.Duration = 50*60 //续租x分钟
	RemainRenewalTime time.Duration = 5*60//剩余x分钟，在下一次请求时开始续租

	GameServerAddress string = "45.77.43.2:6666" //百家乐游戏服务器地址
)


