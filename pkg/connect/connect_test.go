package connect

import "testing"

func TestGet(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "测试能请求通的URL", args: args{url: "http://www.baidu.com"}, want: true},
		{name: "测试不能请求通的URL", args: args{url: "http://invalid.url.test"}, want: false},
		{name: "测试404的URL", args: args{url: "http://www.baidu.com/nonexistentpage"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.url); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
