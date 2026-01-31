package main

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

func TestLRUCache(t *testing.T) {
	tests := []struct {
		name     string // test name
		cmds     string // input arg 1
		args     string
		expected string // expected result
	}{
		// 2. load all test cases
		{"case 1", `["LRUCache","put","put","get","put","get","put","get","get","get"]`,
			`[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]`, `[null,null,null,1,null,-1,null,-1,3,4]`},
	}

	var obj *LRUCache

	mapAction := map[string]func(string) string{
		"LRUCache": func(input string) string {
			var p []int
			json.Unmarshal([]byte(input), &p) // 解析如 "[2]"
			capacity := p[0]
			tmp := Constructor(capacity) // 假設 Constructor 回傳 LRUCache
			obj = &tmp
			return "null"
		},
		"get": func(input string) string {
			var p []int
			json.Unmarshal([]byte(input), &p)
			val := obj.Get(p[0])
			return strconv.Itoa(val)
		},
		"put": func(input string) string {
			var p []int
			json.Unmarshal([]byte(input), &p)
			obj.Put(p[0], p[1])
			return "null"
		},
	}

	for _, tt := range tests {
		var res string
		// 3. use t.Run to create subtests
		t.Run(tt.name, func(t *testing.T) {

			cmds := ParseCmds(tt.cmds)
			args := ParseArgsArray(tt.args)

			var outputs []string
			for i := 0; i < len(cmds); i++ {
				cmd := cmds[i]
				arg := args[i]
				output := mapAction[cmd](arg)
				outputs = append(outputs, output)
			}
			res = SerializeToString(outputs)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("The test %s. The cmds %s with args %s. The result %s; the expected is %s", tt.name, tt.cmds, tt.args, res, tt.expected)
			}
		})
	}
}
