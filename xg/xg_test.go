package xg

import (
	"testing"

	"github.com/pig0224/fjsdxy/test"
)

func TestLogin(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			_, err := Login(tt.args.studentID, tt.args.password)
			if (err != nil) && tt.wantErr {
				t.Errorf("NewCollector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
