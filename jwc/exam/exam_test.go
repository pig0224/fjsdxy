package exam

import (
	"fmt"
	"github.com/pig0224/fjsdxy/jwc"
	"github.com/pig0224/fjsdxy/test"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		studentID string
		password  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		//{
		//	name: "登录失败",
		//	args: args{
		//		studentID: "34567890",
		//		password:  "lalalalalal",
		//	},
		//	wantErr: true,
		//},
		{
			name: "登录成功",
			args: args{
				studentID: test.StudentID,
				password:  test.Password,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run("getSource", func(t *testing.T) {
			c, err := jwc.Login(tt.args.studentID, tt.args.password)
			if err == nil {
				res, err := Get("", c)
				if err == nil {
					fmt.Println(res)
				}
				//fmt.Println(*st)
			}
			if (err != nil) && tt.wantErr {
				t.Errorf("NewCollector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
