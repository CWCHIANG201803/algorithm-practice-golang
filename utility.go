package main

import (
	"encoding/json"
	"fmt"
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

func SerializeToString[T any](arr []T) string {
	if len(arr) == 0 {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")
	for i, v := range arr {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("%v", v))
	}
	sb.WriteString("]")
	return sb.String()
}

func ParseCmds(str string) []string {
	var res []string
	json.Unmarshal([]byte(str), &res)
	return res
}

func ParseArgsArray(jsonStr string) []string {
	var rawParts []json.RawMessage
	err := json.Unmarshal([]byte(jsonStr), &rawParts)
	if err != nil {
		return nil
	}

	// 2. 把每個 RawMessage 轉回字串
	result := make([]string, len(rawParts))
	for i, r := range rawParts {
		result[i] = string(r)
	}
	return result
}

func BuildStringArray(jsonStr string) []string {
	// Remove outer quotes if present
	if strings.HasPrefix(jsonStr, `"`) && strings.HasSuffix(jsonStr, `"`) {
		jsonStr = jsonStr[1 : len(jsonStr)-1]
	}

	var result []string
	var current string
	depth := 0

	for _, char := range jsonStr {
		if char == '[' {
			if depth == 0 {
				current = ""
			}
			current += string(char)
			depth++
		} else if char == ']' {
			depth--
			current += string(char)
			if depth == 0 && current != "" {
				result = append(result, current)
				current = ""
			}
		} else if depth > 0 && char != ' ' {
			current += string(char)
		}
	}

	return result
}
