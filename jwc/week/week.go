package week

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/pig0224/fjsdxy/config"
	"strconv"
	"strings"
	"time"
)

type Week struct {
	Weekly uint64       `json:"weekly"`
	week   time.Weekday `json:"week"`
	Today  string       `json:"today"`
}

func Get(term string, c *colly.Collector) ([]Week, error) {

	var weeks []Week
	var logErr error
	var weekly uint64
	c.OnHTML("#kbtable", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, e *colly.HTMLElement) {
			if i > 0 {
				e.ForEach("td", func(i int, e *colly.HTMLElement) {

					if e.Text != "周历编制" {
						if i == 0 {
							weekly, _ = strconv.ParseUint(e.Text, 10, 64)
						} else if e.ChildText("font") != "" {
							if weekly > 0 {
								var week Week
								week.Weekly = weekly
								week.week = time.Weekday(i - 1)
								week.Today = strings.Replace(e.Attr("title"), "年", "-", 1)
								week.Today = strings.Replace(week.Today, "月", "-", 1)

								weeks = append(weeks, week)
							}
						}
					}
				})
			}
		})
		if len(weeks) <= 0 {
			logErr = errors.New("获取学校周历失败")
		}
	})

	if err := c.Post(config.JW_DOMAIN+"/jsxsd/jxzl/jxzl_query", map[string]string{
		"xnxq01id": term,
	}); err != nil {
		return nil, err
	}

	if logErr != nil {
		return nil, logErr
	}

	return weeks, nil
}
