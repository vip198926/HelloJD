package common

import (
	"github.com/Albert-Zhan/httpc"
	"github.com/unknwon/goconfig"
)

const (
	SoftName          = "jd_seckill"
	Version           = "0.2.8"
	DateTimeFormatStr = "2006-01-02 15:04:05"
	DateFormatStr     = "2006-01-02"
	IniFileContent    = `
[config]
log_level = info
qrcode_show_type = open
qrcode_create_api = https://api.pwmqr.com/qrcode/create/?url=
eid =
fp =
sku_id = 100012043978
seckill_num = 1
buy_time = 2021-01-01 09:59:59
seckill_time = 1
task_num = 1
ticker_time = 
default_user_agent = Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36

[account]
payment_pwd =

[messenger]
enable = false
type = none
email = 
server_chan_sckey =

[smtp]
email_host =
port =
email_user =
email_pwd =

[dingtalk]
access_token =
secret =
at =

[root]
user_name = 
name = 6666
upname = 6666
code = 8888
`
)

var (
	SoftDir string

	Client *httpc.HttpClient

	CookieJar *httpc.CookieJar

	Config *goconfig.ConfigFile

	SeckillStatus chan bool

	ViewQrcodePid = 0
)
