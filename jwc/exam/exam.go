package exam

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/pig0224/fjsdxy/config"
	"strings"
)

type Result struct {
	Name    string `json:"name"`
	Term    string `json:"term"`
	Type    string `json:"type"`
	Score   string `json:"score"`
	Credits string `json:"credits"`
}

func Get(term string, c *colly.Collector) ([]Result, error) {
	var results []Result
	var logErr error

	c.OnHTML("#dataList", func(e *colly.HTMLElement) {

		e.ForEach("tr", func(i int, e *colly.HTMLElement) {
			if i > 0 {
				var result Result
				e.ForEach("td", func(i int, e *colly.HTMLElement) {
					switch i {
					case 1:
						result.Term = e.Text
					case 3:
						result.Name = e.Text
					case 4:
						str := strings.Replace(e.Text, " ", "", -1)
						str = strings.Replace(str, "\n", "", -1)
						str = strings.Replace(str, "\t", "", -1)
						result.Score = str
					case 6:
						result.Credits = e.Text
					case 12:
						result.Type = e.Text
					}
				})
				results = append(results, result)
			}
		})
		if len(results) <= 0 {
			logErr = errors.New("获取成绩失败")
		}
	})

	if err := c.Post(config.JW_DOMAIN+"/jsxsd/kscj/cjcx_list", map[string]string{
		"kksj": term,
		"xsfs": "all",
	}); err != nil {
		logErr = errors.New("获取成绩失败")
	}

	if logErr != nil {
		return nil, logErr
	}

	return results, nil
}
