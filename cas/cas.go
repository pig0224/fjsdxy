package cas

import (
	"errors"
	"github.com/gocolly/colly"
	sy "github.com/pig0224/fjsdxy"
	"github.com/pig0224/fjsdxy/config"
)

// Login 登录CAS，获取已登录采集器
func Login(studentID, password string) (*colly.Collector, error) {
	c, login_err := sy.SSO_Login(studentID, password)

	// 判定是否登录失败
	if login_err != nil {
		return nil, errors.New("登陆失败，请检查用户名密码是否正确！")
	}

	// 尝试登录
	if err := c.Visit(config.CAS_DOMAIN + "/cas/login?mode=rlogin&service=http://cas.fjsdxy.com/portal/a/shiro-cas"); err != nil {
		return nil, err
	}

	return c, nil
}

// Logout 退出CAS系统
func Logout(c *colly.Collector) error {
	return c.Visit(config.CAS_DOMAIN + "/portal/a/logout")
}
