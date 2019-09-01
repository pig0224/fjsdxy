package leave

import (
	"github.com/pig0224/fjsdxy/xg"
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
				studentID: "1627160237",
				password:  "080025",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run("getSource", func(t *testing.T) {
			c, err := xg.Login(tt.args.studentID, tt.args.password)
			if err == nil {
				Apply(c)
				//fmt.Println(*st)
			}
			if (err != nil) && tt.wantErr {
				t.Errorf("NewCollector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
