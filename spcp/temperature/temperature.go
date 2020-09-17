package temperature

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/pig0224/fjsdxy/config"
	"strings"
)

type FillForm struct {
	TimeNowHour   string
	TimeNowMinute string
	Temper1       string
	Temper2       string
	ReSubmiteFlag string
}

func Fill(c *colly.Collector, TimeNowHour, TimeNowMinute, Temper1, Temper2 string) (*colly.Collector, error) {
	var logErr error
	fillForm := FillForm{TimeNowHour: TimeNowHour, TimeNowMinute: TimeNowMinute, Temper1: Temper1, Temper2: Temper2}
	fillForm = GetFillForm(c, fillForm)

	if fillForm.ReSubmiteFlag == "" {
		logErr = errors.New("系统出现故障")
	}

	c.OnResponse(func(r *colly.Response) {
		//判断是否填报成功
		res := string(r.Body)
		//println(res)
		if !strings.Contains(res, "填报成功") {
			logErr = errors.New("填报失败")
		}
	})

	err := c.Post(config.XG_DOMAIN+"/SPCP/Web/Temperature/StuTemperatureInfo", map[string]string{
		"TimeNowHour":   fillForm.TimeNowHour,
		"TimeNowMinute": fillForm.TimeNowMinute,
		"Temper1":       fillForm.Temper1,
		"Temper2":       fillForm.Temper2,
		"ReSubmiteFlag": fillForm.ReSubmiteFlag,
	})

	if err != nil {
		return nil, err
	}

	if logErr != nil {
		return c, logErr
	}

	return c, nil
}

func GetFillForm(c *colly.Collector, form FillForm) FillForm {
	c.OnHTML("input", func(e *colly.HTMLElement) {
		switch e.Attr("name") {
		case "ReSubmiteFlag":
			form.ReSubmiteFlag = e.Attr("value")
		}
	})

	_ = c.Visit(config.XG_DOMAIN + "/SPCP/Web/Temperature/StuTemperatureInfo")

	return form
}
