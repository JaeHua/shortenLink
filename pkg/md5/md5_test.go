package md5

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "测试空字符串", args: args{data: []byte("")}, want: "d41d8cd98f00b204e9800998ecf8427e"},
		{name: "测试普通字符串", args: args{data: []byte("hello world")}, want: "5eb63bbbe01eeed093cb22bb8f5acdc3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.data); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
