package ecard

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/pig0224/fjsdxy/config"
)

type Ecard struct {
	Code    string `json:"code"`
	Money   string `json:"money"`
	Consume string `json:"consume"`
}

func Get(c *colly.Collector) (*Ecard, error) {
	var ecard Ecard
	var logErr error

	c.OnHTML(".personal_content", func(e *colly.HTMLElement) {
		e.ForEach("tbody", func(i int, e *colly.HTMLElement) {
			if i == 0 {
				e.ForEach("td", func(i int, e *colly.HTMLElement) {
					switch i {
					case 0:
						ecard.Code = e.Text
					case 1:
						ecard.Money = e.Text
					case 2:
						ecard.Consume = e.Text
					}
				})
			}
		})

		if ecard.Code == "" {
			logErr = errors.New("获取一卡通信息失败")
		}

	})

	if err := c.Visit(config.CAS_DOMAIN + "/portal/a/myEkt"); err != nil {
		return nil, err
	}

	if logErr != nil {
		return nil, logErr
	}

	return &ecard, nil
}
