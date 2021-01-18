package cmd

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/unknwon/goconfig"
	"github.com/ztino/jd_seckill/common"
	"github.com/ztino/jd_seckill/jd_seckill"
	"github.com/ztino/jd_seckill/log"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Open JD’s simulated login",
	Run:   startLogin,
}

func startLogin(cmd *cobra.Command, args []string) {

	var name string

	session := jd_seckill.NewSession(common.CookieJar)
	//检测是否登录过
	if common.Exists("./cookie.txt") {
		//已登录，检测登录状态
		err := session.CheckLoginStatus()
		if err != nil {
			log.Error("登录失效，请重新登录")
			return
		}
		user := jd_seckill.NewUser(common.Client, common.Config)
		log.Warn("登录成功")
		userInfo, _ := user.GetUserInfo()
		name = userInfo
		log.Warn("用户:" + userInfo)

	} else {
		//未登录
		user := jd_seckill.NewUser(common.Client, common.Config)
		wlfstkSmdl, err := user.QrLogin()
		defer user.DelQrCode()
		if err != nil {
			os.Exit(0)
		}
		ticket := ""
		for {
			ticket, err = user.QrcodeTicket(wlfstkSmdl)
			if err == nil && ticket != "" {
				break
			}
			time.Sleep(2 * time.Second)
		}
		_, err = user.TicketInfo(ticket)
		if err == nil {
			if status := user.RefreshStatus(); status == nil {
				//保存cookie
				_ = session.SaveCookieToFile("./cookie.txt")
				log.Warn("登录成功")
				userInfo, _ := user.GetUserInfo()
				name = userInfo
				log.Warn("用户:" + userInfo)
			} else {
				log.Error("登录失效")
			}
		} else {
			log.Error("登录失败")
		}
	}
	//保存用户名到配置文件中
	confFile := common.SoftDir + "/conf.ini"
	cfg, err := goconfig.LoadConfigFile(confFile)
	if err != nil {
		log.Error("配置出错，程序退出，请重新启动")
		os.Exit(0)
	}
	//把用户名name保存到配置文件[root]中，实现外部调用（易语言）
	cfg.SetValue("root", "user_name", name)
	if err := goconfig.SaveConfigFile(cfg, confFile); err != nil {
		log.Error("配置失败，请重新启动", name)
	}
}
