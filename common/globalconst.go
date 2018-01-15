package common

import "time"

const (
	RedisStorageTime time.Duration = 20*60 //默认redis存储时间x分钟
	RenewalTime time.Duration = 25*60 //续租x分钟
	RemainRenewalTime time.Duration = 5*60//剩余x分钟，在下一次请求时开始续租
)
