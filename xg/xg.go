package xg

import (
	"errors"
	sy "fjsdxy"
	"fjsdxy/config"

	"github.com/gocolly/colly"
)

// Login 登录教务处，获取已登录采集器
func Login(studentID, password string) (*colly.Collector, error) {
	c, _ := sy.SSO_Login(studentID, password)
	// 判定是否登录失败
	var logErr error
	c.OnHTML("#fm1", func(e *colly.HTMLElement) {
		logErr = errors.New("登陆失败，请检查用户名密码是否正确！")
	})

	// 尝试登录
	if err := c.Visit(config.CAS_DOMAIN + "?service=" + config.XG_DOMAIN + "/Sys/LoginOne.aspx"); err != nil {
		return nil, err
	}

	// 如果登录失败，返回登录信息
	if logErr != nil {
		return nil, logErr
	}
	return c, nil
}

// Logout 退出教务处
func Logout(c *colly.Collector) error {
	return c.Visit(config.XG_DOMAIN + "/Sys/SystemForm/ExitWindows.aspx")
}
