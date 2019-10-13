package course

import (
	"encoding/json"
	"errors"
	"github.com/gocolly/colly"
	"github.com/pig0224/fjsdxy/config"
	"strconv"
)

type Course struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Class    string `json:"class"`
	Teacher  string `json:"teacher"`
}

func Get(term string, week int, c *colly.Collector) ([]Course, error) {
	var courses []Course
	var logErr error

	res := struct {
		Kbxx [][]string `json:"kbxx"`
	}{}

	c.OnResponse(func(r *colly.Response) {
		_ = json.Unmarshal(r.Body, &res)
		for i, v := range res.Kbxx {
			if i > 0 {
				var course Course
				course.Name = v[0]
				course.Class = v[4]
				course.Position = v[1]
				course.Teacher = v[6]
				courses = append(courses, course)
			}
		}
		if len(courses) <= 0 {
			logErr = errors.New("课表获取失败")
		}

	})

	if err := c.Visit(config.CAS_DOMAIN + "/portal/a/getStTimetableByweek?xnxq=" + term + "&week=" + string(strconv.Itoa(int(week)))); err != nil {
		return nil, err
	}

	if logErr != nil {
		return nil, logErr
	}

	return courses, nil
}
