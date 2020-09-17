package spcp

import (
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
			//c, err := AutoLogin("E385925CA3A40E02E87FE1A7192FD80297E5A316202AA853F3D05D7E6FE955124FF7435563208A5FE5FCBFDCC9E3DB7FDA322AED2631B0C20CC5725FACE6CE29E0A29387591E4F1D60D5ACA3926D0FE3991E7B7617A0DCF38F80E077EEEE08BE13438D7C076C2A1A5AEECEC98A917F75542FF014F0A33612B47D6390A5BC0D5338728D2C2C437B8E5F93581AAEB4090A25FE491F167D3E230649722A3DC5D7778F110132C66C937BF58AC459F0A49BB3EF67854E6FECDEF22D3434EECB5CFA338D21A287C49BB9A32FE426B38A2960C3C02BEBA89DEAA8BC85D26FAA126793C58FA316F24C1804485C3B5CF599228C33D2D4A5B34595AD64499969A382FD30FF")
			//if err == nil {
			//	_, err = temperature.Fill(c, "12", "30", "36", "5")
			//	println(err == nil)
			//}
		})
	}
}
