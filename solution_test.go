package solution

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string // test name
		input    string // input arg
		expected string // expected result
	}{
		// 2. load all test cases
		{"case 1", "[1,2,3,4,5]", "[5,4,3,2,1]"},
		{"case 2", "[]", "[]"},
		{"case 3", "[1]", "[1]"},
		{"case 4", "[1,2]", "[2,1]"},
	}

	for _, tt := range tests {
		// 3. use t.Run to create subtests
		t.Run(tt.name, func(t *testing.T) {

			lst := BuildLinkedList(tt.input)
			got := SerializeLinkedListToStr(reverseList(lst))

			// 4. validate result
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("reverseList(%v). Got result %v; the expected is %v", tt.input, got, tt.expected)
			}
		})
	}
}
