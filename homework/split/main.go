package split

import (
	"strings"
)

func Splitstr(s, sep string) []string {
	//可以先配置容量 这样append就不会出错 ，如果直接申请长度会出现多余的0
	result := make([]string, 0, strings.Count(s, sep)+1)
	if sep == "" {
		for _, v := range s {
			result = append(result, string(v))
		}
		return result
	} else if strings.Contains(s, sep) {
		length := len(sep)
		start := 0
		temp := s
		for {
			index := strings.Index(temp, sep)
			if index == -1 {
				result = append(result, temp)
				break
			}
			str := temp[0:index]
			if len(str) != 0 {
				result = append(result, str)
			}
			start = index + length
			if start >= len(temp) {
				break
			}
			temp = temp[start:]
		}
		return result
	}
	return nil
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
