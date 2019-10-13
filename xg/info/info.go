package info

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/pig0224/fjsdxy/config"
)

type XgInfo struct {
	AvatorImg string
	StudentId string
	Password  string
}

func GetInfo(studentID string, password string, c *colly.Collector) (*XgInfo, error) {
	var logErr error

	xgInfo := XgInfo{}
	c.OnHTML("#Student11_Image1", func(e *colly.HTMLElement) {
		if e.Attr("src") != "" {
			xgInfo.AvatorImg = "http://xg.fjsdxy.com/Sys/SystemForm/Class/" + e.Attr("src")
			xgInfo.StudentId = studentID
			xgInfo.Password = password
		} else {
			logErr = errors.New("获取头像失败")
		}
	})
	if err := c.Visit(config.XG_DOMAIN + "/Sys/SystemForm/Class/MyStudent.aspx"); err != nil {
		return nil, err
	}

	if logErr != nil {
		return nil, logErr
	}

	return &xgInfo, nil
}
