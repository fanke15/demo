package pkg

import (
	"cron/internal/lib/define"
	"strconv"
	"strings"
)

// AnyToStr 将任意类型转换为string类型
func AnyToStr(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch vt := value.(type) {
	case float64:
		key = strconv.FormatFloat(vt, 'f', define.NegativeOne, define.SixtyFour)
	case float32:
		key = strconv.FormatFloat(float64(vt), 'f', define.NegativeOne, define.SixtyFour)
	case int:
		key = strconv.Itoa(vt)
	case uint:
		key = strconv.Itoa(int(vt))
	case int8:
		key = strconv.Itoa(int(vt))
	case uint8:
		key = strconv.Itoa(int(vt))
	case int16:
		key = strconv.Itoa(int(vt))
	case uint16:
		key = strconv.Itoa(int(vt))
	case int32:
		key = strconv.Itoa(int(vt))
	case uint32:
		key = strconv.Itoa(int(vt))
	case int64:
		key = strconv.FormatInt(vt, define.Ten)
	case uint64:
		key = strconv.FormatUint(vt, define.Ten)
	case string:
		key = vt
	case []byte:
		key = string(vt)
	default:
		key = string(MarshalToString(value))
	}
	return key
}

// AnySliceToStr 字符串拼接
func AnySliceToStr(sep string, strs ...string) string {
	var (
		build strings.Builder
		total = len(strs)
	)
	for k, v := range strs {
		build.WriteString(v)
		if total-define.One > k {
			build.WriteString(sep)
		}
	}
	return build.String()
}
