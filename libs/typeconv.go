package libs

import (
	"strconv"
	"strings"
)

func split(s string) []string {
	ret := make([]string, len([]rune(s)))
	for i, v := range []rune(s) {
		ret[i] = string(v)
	}
	return ret
}

func join(x []string) string {
	return strings.Join(x, "")
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func intToStr(i int) string {
	return strconv.Itoa(i)
}
