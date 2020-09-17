package spcp

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/gocolly/colly"
	"github.com/pig0224/fjsdxy/config"
	"io/ioutil"
	"strings"
)

type LoginForm struct {
	StudentId     string `json:"txtUid"`
	Password      string `json:"txtPwd"`
	Code          string `json:"code"`
	ReSubmiteFlag string
	StuLoginMode  string
}

// Cookie自动登录
func AutoLogin(cookie string) (*colly.Collector, error) {
	var c = colly.NewCollector()
	var logErr error
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", "CenterSoftWeb="+cookie)
	})
	c.OnResponse(func(r *colly.Response) {
		//判断是否填报成功
		res := string(r.Body)
		if !strings.Contains(res, "安全退出") {
			logErr = errors.New("登录失败")
		}
	})
	err := c.Visit(config.XG_DOMAIN + "/SPCP/Web/")
	if err != nil {
		return nil, err
	}

	if logErr != nil {
		return nil, logErr
	}
	return c, nil
}

// 登录填报系统
func Login(studentID, password, code, codeCookie string) (*colly.Collector, string, error) {
	var c = colly.NewCollector()
	var logErr error
	var CenterSoftWeb = ""
	loginForm := LoginForm{StudentId: studentID, Password: password, Code: code}
	loginForm = GetLoginForm(c, loginForm)

	if loginForm.ReSubmiteFlag == "" {
		logErr = errors.New("系统出现故障")
	}

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", "ASP.NET_SessionId="+codeCookie)
	})

	c.OnResponse(func(r *colly.Response) {
		//res := string(r.Body)
		//println(res)
		siteCokkie := c.Cookies(config.XG_DOMAIN + "/SPCP/Web/")
		if len(siteCokkie) > 0 {
			for _, cc := range siteCokkie {
				if cc.Name == "CenterSoftWeb" {
					CenterSoftWeb = cc.Value
					logErr = nil
				}
			}
		} else {
			logErr = errors.New("登录失败")
		}

	})

	err := c.Post(config.XG_DOMAIN+"/SPCP/Web/", map[string]string{
		"ReSubmiteFlag": loginForm.ReSubmiteFlag,
		"StuLoginMode":  loginForm.StuLoginMode,
		"txtUid":        loginForm.StudentId,
		"txtPwd":        loginForm.Password,
		"code":          loginForm.Code,
	})

	if err != nil {
		return nil, CenterSoftWeb, err
	}

	if logErr != nil {
		return nil, CenterSoftWeb, errors.New("登录失败")
	}

	return c, CenterSoftWeb, nil
}

// 退出填报系统
func Logout() error {
	var c = colly.NewCollector()
	return c.Visit(config.XG_DOMAIN + "/SPCP/Web/Account/Logout")
}

//获取登录表单
func GetLoginForm(c *colly.Collector, loginForm LoginForm) LoginForm {

	c.OnHTML("input", func(e *colly.HTMLElement) {
		switch e.Attr("name") {
		case "ReSubmiteFlag":
			loginForm.ReSubmiteFlag = e.Attr("value")
		case "StuLoginMode":
			loginForm.StuLoginMode = e.Attr("value")
		}
	})

	_ = c.Visit(config.XG_DOMAIN + "/SPCP/Web/")

	return loginForm
}

//获取验证码
func GetLoginCode() (base64Img string, codeCookies string, err error) {
	var c = colly.NewCollector()
	var logErr error
	base64Img = ""
	codeCookies = ""
	c.OnResponse(func(r *colly.Response) {
		srcByte, err := ioutil.ReadAll(bytes.NewReader(r.Body))
		if err != nil {
			logErr = errors.New("验证码获取失败")
		}
		res := base64.StdEncoding.EncodeToString(srcByte)
		base64Img = "data:image/png;base64," + res
		siteCokkie := c.Cookies(config.XG_DOMAIN + "/SPCP/Web/")
		for _, cc := range siteCokkie {
			if cc.Name == "ASP.NET_SessionId" {
				codeCookies = cc.Value
			}
		}
	})

	_ = c.Visit(config.XG_DOMAIN + "/SPCP/Web/Account/GetLoginVCode")

	if logErr != nil {
		return "", "", logErr
	}
	if codeCookies == "" || base64Img == "" {
		logErr = errors.New("验证码获取失败")
		return "", "", logErr
	}
	return base64Img, codeCookies, nil
}
