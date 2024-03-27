package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ArrayIntToString(input []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(input), " ", delim, -1), "[]")
}

func StringToIntArray(s string) []int {
	if s == "" {
		return make([]int, 0)
	}
	arr := strings.Split(s, ",")
	b := make([]int, len(arr))
	for i, v := range arr {
		b[i], _ = strconv.Atoi(v)
	}
	return b
}
