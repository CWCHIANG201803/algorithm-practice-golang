package solution

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string // test name
		nums     string // input arg 1
		target   int    // input arg 2
		expected string // expected result
	}{
		// 2. load all test cases
		{"case 1", "[2,7,11,15]", 9, "[0,1]"},
		{"case 2", "[3,2,4]", 6, "[1,2]"},
		{"case 3", "[3,3]", 6, "[0,1]"},
	}

	for _, tt := range tests {
		// 3. use t.Run to create subtests
		t.Run(tt.name, func(t *testing.T) {
			nums := BuildArray(tt.nums)
			got := twoSum(nums, tt.target)
			res := SerialzeToString(got)
			// 4. validate result
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("twoSum(%s, %d) = %s; the expected is %s", tt.nums, tt.target, res, tt.expected)
			}
		})
	}
}
