package base62

import (
	"math"
	"strings"
)

// 62 进制转换模块
//012346789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ

var (
	base62String string
)

func MustInit(bs string) {
	if len(bs) == 0 {
		panic("base62 string can not be empty")
	}
	base62String = bs
}

// 为了避免被恶意请求，可以对base62String进行打乱

// Int2String 10进制数字转62进制字符串
func Int2String(seq uint64) string {
	if seq == 0 {
		return string(base62String[0])
	}
	result := make([]byte, 0)
	base := uint64(len(base62String))
	for seq > 0 {
		remainder := seq % base
		result = append([]byte{base62String[remainder]}, result...)
		seq = seq / base
	}
	return string(result)
}

// String2Int 62进制字符串转10进制数字
func String2Int(s string) uint64 {
	bl := []byte(s)
	// 反转bl
	l := len(bl)
	for i := 0; i < l/2; i++ {
		bl[i], bl[l-1-i] = bl[l-1-i], bl[i]
	}
	var result uint64 = 0
	for idx, b := range bl {
		base := math.Pow(62, float64(idx))
		result += uint64(strings.Index(base62String, string(b))) * uint64(base)
	}
	return result
}
