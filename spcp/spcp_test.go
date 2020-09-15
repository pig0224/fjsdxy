package spcp

import (
	"github.com/pig0224/fjsdxy/spcp/temperature"
	"testing"
)

func TestSSOLogin(t *testing.T) {
	type args struct {
		studentID string
		password  string
		code      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "登录测试",
			args: args{
				studentID: "1380180058",
				password:  "244538",
				code:      "sfdf",
			},
			wantErr: true,
		},
		//{
		//	name: "登录成功",
		//	args: args{
		//		studentID: test.StudentID,
		//		password:  test.Password,
		//	},
		//	wantErr: true,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			////img, _  := GetLoginCode()
			////println(img)
			//_, err := Login(tt.args.studentID, tt.args.password, tt.args.code)
			//println(CenterSoftWeb)
			//
			//if (err != nil) && tt.wantErr {
			//	t.Errorf("NewCollector() error = %v, wantErr %v", err, tt.wantErr)
			//	return
			//}
			c, err := AutoLogin("1D6C64122640905AD85B45774A9C82DD655C0C81F72E9E03756A87EF53EA464FB882D2A1B313DD2B657D9FD96814260A9762B51C50A618142F8AE11F979597200D0B33090EA6A9615ACC1B932B8B1DE7F55B6416A50212CBC8322C068247F815EC2BC946574EC47A9690BDDC727E31D5130071B58887B75B5E108C939B23F15751C80DB57947630B1508F768686E2C9DADEC62735B516E886951EE18C3BB3194090DAF1F72435DE4F69F23313E14FA4C71318B85EB9F817C19C517EA10D3F5FAC0D547ED8806865DBB72B86E1658F4E0ED679FA0D826287F4CF185EC22BD1360122E724018DAFBB37DE0D85867B45E709119FDA34174AB85185D8182CBFCADF7")
			if err == nil {
				_, err = temperature.Fill(c, "12", "30", "36", "5")
				println(err == nil)
			}
		})
	}
}
