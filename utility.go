package solution

import (
	"strconv"
	"strings"
)

func BuildArray(s string) []int {
	s = strings.Trim(s, "[]")
	elements := strings.Split(s, ",")

	result := make([]int, len(elements))
	for i, elem := range elements {
		num, err := strconv.Atoi(strings.TrimSpace(elem))

		if err != nil {
			return nil
		}
		result[i] = num
	}

	return result
}

func SerialzeToString(arr []int) string {
	if len(arr) == 0 {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")
	for i, num := range arr {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(strconv.Itoa(num))
	}
	sb.WriteString("]")
	return sb.String()
}
