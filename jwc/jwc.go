package jwc

import (
	"errors"
	sy "fjsdxy"
	"fjsdxy/config"

	"github.com/gocolly/colly"
)

// Login 登录教务处，获取已登录采集器
func Login(studentID, password string) (*colly.Collector, error) {
	c, login_err := sy.SSO_Login(studentID, password)

	// 判定是否登录失败
	if login_err != nil {
		return nil, errors.New("登陆失败，请检查用户名密码是否正确！")
	}

	// 尝试登录
	if err := c.Visit(config.CAS_DOMAIN + "?service=" + config.JW_DOMAIN + "/jsxsd/sso.jsp"); err != nil {
		return nil, err
	}

	return c, nil
}

// Logout 退出教务处
func Logout(c *colly.Collector) error {
	return c.Visit(config.JW_DOMAIN + "/Logon.do?method=logoutFromJsxsd")
}
