package urltool

import "testing"

func TestGetBasePath(t *testing.T) {
	type args struct {
		targetUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "基本测试", args: args{targetUrl: "http://example.com/path/to/resource"}, want: "resource", wantErr: false},
		{name: "带查询参数", args: args{targetUrl: "http://example.com/path/to/resource?query=123"}, want: "resource", wantErr: false},
		{name: "根路径", args: args{targetUrl: "http://example.com/"}, want: "", wantErr: true},
		{name: "无路径", args: args{targetUrl: "http://example.com"}, want: "", wantErr: true},
		{name: "无效URL", args: args{targetUrl: "/xxxx/1123"}, want: "", wantErr: true},
		{name: "路径以斜杠结尾", args: args{targetUrl: "http://example.com/path/to/"}, want: "to", wantErr: false},
		{name: "空字符串", args: args{targetUrl: ""}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBasePath(tt.args.targetUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBasePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBasePath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
