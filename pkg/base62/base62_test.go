package base62

import "testing"

func TestInt2String(t *testing.T) {
	type args struct {
		seq uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test0", args{0}, "0"},
		{"test1", args{1}, "1"},
		{"test61", args{61}, "Z"},
		{"test62", args{62}, "10"},
		{"test3843", args{3843}, "ZZ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int2String(tt.args.seq); got != tt.want {
				t.Errorf("Int2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString2Int(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
		{"test0", args{"0"}, 0},
		{"test1", args{"1"}, 1},
		{"test61", args{"Z"}, 61},
		{"test62", args{"10"}, 62},
		{"test3843", args{"ZZ"}, 3843},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String2Int(tt.args.s); got != tt.want {
				t.Errorf("String2Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
