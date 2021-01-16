package service

import (
	"github.com/unknwon/goconfig"
)

func SendMessage(conf *goconfig.ConfigFile, title, msg string) error {
	// 收件方
	mailTo := []string{
		"qsyxcx@qq.com", // 这里最好写成邮箱收发件时的这种标记格式
	}
	email := NewEmail(conf)
	err := email.Send(mailTo, title, msg)
	// log.Info(err)
	return err
	// if conf.MustValue("messenger", "enable", "false") == "true" {
	// 	//钉钉机器人
	// 	if conf.MustValue("messenger", "type", "none") == "dingtalk" {
	// 		dingtalk := NewDingtalk(conf)
	// 		err := dingtalk.Send(title, msg)
	// 		return err
	// 	}
	// 	//邮件发送
	// 	if conf.MustValue("messenger", "type", "none") == "smtp" {
	// 		email := NewEmail(conf)
	// 		err := email.Send([]string{conf.MustValue("messenger", "email", "")}, title, msg)
	// 		return err
	// 	}
	// 	//Server酱推送
	// 	if conf.MustValue("messenger", "type", "none") == "wechat" {
	// 		wechat := NewWechat(conf)
	// 		err := wechat.Send(title, msg)
	// 		return err
	// 	}
	// }
	// return nil
}
