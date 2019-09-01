package leave

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/pig0224/fjsdxy/config"
	"strconv"
	"strings"
)

type Leave struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	Date     string `json:"date"`
	AreaWide string `json:"AreaWide"`
	Status   string `json:"status"`
}

type LeaveInfo struct {
	LeaveBeginDate string `json:"leave_begin_date"`
	LeaveBeginTime string `json:"leave_begin_time"`
	LeaveEndDate   string `json:"leave_end_date"`
	LeaveEndTime   string `json:"leave_end_time"`
	LeaveType      string `json:"leave_type"`
	OutAddress     string `json:"out_address"`
	AreaWide       string `json:"area_wide"`
	OutMoveTel     string `json:"out_move_tel"`
	Relation       string `json:"relation"`
	OutName        string `json:"out_name"`
	StuMoveTel     string `json:"stu_move_tel"`
	LeaveThing     string `json:"leave_thing"`
}

func Get(c *colly.Collector) ([]Leave, error) {
	var leaveList []Leave
	var logErr error

	c.OnHTML("#GridView1", func(e *colly.HTMLElement) {
		e.ForEach("tbody", func(i int, e *colly.HTMLElement) {
			if i == 0 {
				e.ForEach("tr", func(i int, e *colly.HTMLElement) {
					if i != 0 {
						var leave Leave
						e.ForEach("td", func(i int, e *colly.HTMLElement) {
							switch i {
							case 1: //Id
								idUrl := e.ChildAttr("a", "href")
								Id := idUrl[43:len(idUrl)]
								leave.Id, _ = strconv.Atoi(Id)
							case 2: //时间
								date := strings.Replace(e.Text, "年", ".", -1)
								date = strings.Replace(date, "月", ".", -1)
								date = strings.Replace(date, "日", ".", -1)
								date = strings.Replace(date, "点", "", -1)
								date = strings.Replace(date, "至", " - ", -1)
								leave.Date = date
							case 3: //类型
								leave.Type = e.Text
							case 4: //地点
								leave.AreaWide = e.Text
							case 6: //状态
								leave.Status = e.Text
							}
						})
						leaveList = append(leaveList, leave)
					}
				})
			}
		})
		if len(leaveList) == 0 {
			logErr = errors.New("无请假数据")
		}
	})

	if err := c.Visit(config.XG_DOMAIN + "/Sys/SystemForm/Leave/StuAllLeaveManage.aspx"); err != nil {
		return nil, err
	}

	if logErr != nil {
		return nil, logErr
	}

	return leaveList, nil
}

func Revoke(Id int, c *colly.Collector) error {

	var logErr error
	c.OnResponse(func(r *colly.Response) {
		res := string(r.Body)
		if !strings.Contains(res, "操作成功") {
			logErr = errors.New("撤销失败")

		}
	})
	id := strconv.Itoa(Id)
	if err := c.PostMultipart(config.XG_DOMAIN+"/Sys/SystemForm/Leave/StuAllLeaveManage_Edit.aspx?Status=Edit&Id="+id, map[string][]byte{
		"__EVENTTARGET": []byte("Del"),
	}); err != nil {
		return err
	}

	if logErr != nil {
		return logErr
	}

	return nil
}

func Apply(leaveInfo LeaveInfo, c *colly.Collector) error {
	var logErr error
	c.OnResponse(func(r *colly.Response) {
		res := string(r.Body)
		if !strings.Contains(res, "操作成功") {
			logErr = errors.New("申请失败")

		}
	})
	if err := c.PostMultipart(config.XG_DOMAIN+"/Sys/SystemForm/Leave/StuAllLeaveManage_Edit.aspx?Status=Add", map[string][]byte{
		"__EVENTTARGET":            []byte("Save"),
		"AllLeave1$LeaveBeginDate": []byte(leaveInfo.LeaveBeginDate),
		"AllLeave1$LeaveBeginTime": []byte(leaveInfo.LeaveBeginTime),
		"AllLeave1$LeaveEndDate":   []byte(leaveInfo.LeaveEndDate),
		"AllLeave1$LeaveEndTime":   []byte(leaveInfo.LeaveEndTime),
		"AllLeave1$LeaveType":      []byte(leaveInfo.LeaveType),
		"AllLeave1$OutAddress":     []byte(leaveInfo.OutAddress),
		"AllLeave1$AreaWide":       []byte(leaveInfo.AreaWide),
		"AllLeave1$OutMoveTel":     []byte(leaveInfo.OutMoveTel),
		"AllLeave1$Relation":       []byte(leaveInfo.Relation),
		"AllLeave1$OutName":        []byte(leaveInfo.OutName),
		"AllLeave1$StuMoveTel":     []byte(leaveInfo.StuMoveTel),
	}); err != nil {
		return err
	}
	if logErr != nil {
		return logErr
	}
	return nil
}
