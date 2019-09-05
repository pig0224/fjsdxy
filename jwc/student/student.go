package student

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/pig0224/fjsdxy/config"
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

	c.OnHTML(".account ul ", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, e *colly.HTMLElement) {
			switch i {
			case 0:
				student.StudentName = e.Text[12 : len(e.Text)-3]
			case 1:
				student.College = e.ChildText("span")
			case 2:
				student.Major = e.ChildText("span")
			case 3:
				student.ClassName = e.ChildText("span")
			}
		})
		if student.StudentName == "" {
			logErr = errors.New("获取信息失败")
		}
	})
	if err := c.Visit(config.CAS_DOMAIN + "/portal/a"); err != nil {
		return nil, err
	}

	if logErr != nil {
		return nil, logErr
	}

	return &student, nil

}
