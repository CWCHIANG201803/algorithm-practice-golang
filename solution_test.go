package solution

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string // test name
		nums     []int  // input arg 1
		target   int    // input arg 2
		expected []int  // expected result
	}{
		// 2. load all test cases
		{"case 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"case 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"case 3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		// 3. use t.Run to create subtests
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)

			// 4. validate result
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("twoSum(%v, %d) = %v; the expected is %v", tt.nums, tt.target, got, tt.expected)
			}
		})
	}
}
