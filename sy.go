package sy

import (
	"errors"
	"fjsdxy/config"
	"github.com/gocolly/colly"
)

type LoginForm struct {
	studentId  string `json:"username"`
	password   string `json:"password"`
	lt         string
	execution  string
	login_from string
	_eventId   string
}

// 登录统一认证平台
func SSO_Login(studentID, password string) (*colly.Collector, error) {
	var c = colly.NewCollector()
	var logErr error

	loginForm := LoginForm{studentId: studentID, password: password}
	loginForm = GetLoginForm(c, loginForm)

	if loginForm.execution == "" {
		logErr = errors.New("系统出现故障")
	}

	c.OnHTML("#msg", func(e *colly.HTMLElement) {
		logErr = errors.New(e.Text)
	})
	c.OnResponse(func(response *colly.Response) {

	})
	err := c.Post(config.CAS_DOMAIN+"/cas/login", map[string]string{
		"username":   loginForm.studentId,
		"password":   loginForm.password,
		"lt":         loginForm.lt,
		"execution":  loginForm.execution,
		"login_from": loginForm.login_from,
		"_eventId":   loginForm._eventId,
	})

	if err != nil {
		return nil, err
	}

	if logErr != nil {
		return nil, logErr
	}

	return c, nil
}

//获取登录表单
func GetLoginForm(c *colly.Collector, loginForm LoginForm) LoginForm {

	c.OnHTML("input", func(e *colly.HTMLElement) {
		switch e.Attr("name") {
		case "lt":
			loginForm.lt = e.Attr("value")
		case "execution":
			loginForm.execution = e.Attr("value")
		case "login_from":
			loginForm.login_from = e.Attr("value")
		case "_eventId":
			loginForm._eventId = e.Attr("value")
		}
	})

	c.Visit(config.CAS_DOMAIN + "/cas/login")

	return loginForm
}
