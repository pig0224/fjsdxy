package student

import (
	"errors"
	"fjsdxy/config"
	"github.com/gocolly/colly"
)

type Student struct {
	StudentName string `json:"student_name"`
	ClassName   string `json:"class_name"`
	College     string `json:"college"`
	Major       string `json:"major"`
}

func Get(c *colly.Collector) (info *Student, err error) {
	var student Student
	var logErr error

	c.OnHTML(".middletopttxlr", func(e *colly.HTMLElement) {
		e.ForEach(".middletopdwxxcont", func(i int, e *colly.HTMLElement) {
			switch i {
			case 1:
				student.StudentName = e.Text
			case 3:
				student.College = e.Text
			case 4:
				student.Major = e.Text
			case 5:
				student.ClassName = e.Text
			}
		})
		if student.StudentName == "" {
			logErr = errors.New("获取信息失败")
		}
	})
	if err := c.Visit(config.JW_DOMAIN + "/jsxsd/framework/xsMain_new.jsp"); err != nil {
		return nil, err
	}

	if logErr != nil {
		return nil, logErr
	}

	return &student, nil

}
