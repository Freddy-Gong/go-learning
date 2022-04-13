package split

import (
	"strings"
)

func Splitstr(s, sep string) []string {
	result := make([]string, 0)
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
