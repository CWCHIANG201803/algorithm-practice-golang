package solution

import (
	"strings"
	"sync"
	"testing"
)

func TestPrintInOrder(t *testing.T) {
	tests := []struct {
		name     string // test name
		inputs   string // input arg 1
		expected string // expected result
	}{
		// 2. load all test cases
		{"case 1", "[1,2,3]", "firstsecondthird"},
		{"case 2", "[2,1,3]", "firstsecondthird"},
		{"case 3", "[1,3,2]", "firstsecondthird"},
		{"case 4", "[3,2,1]", "firstsecondthird"},
	}

	for _, tt := range tests {
		// 3. use t.Run to create subtests
		t.Run(tt.name, func(t *testing.T) {
			var mu sync.Mutex
			var actualOutput []string
			var wg sync.WaitGroup

			testPrintFirst := func() {
				mu.Lock()
				actualOutput = append(actualOutput, "first")
				mu.Unlock()
			}
			testPrintSecond := func() {
				mu.Lock()
				actualOutput = append(actualOutput, "second")
				mu.Unlock()
			}
			testPrintThird := func() {
				mu.Lock()
				actualOutput = append(actualOutput, "third")
				mu.Unlock()
			}

			foo := NewFoo()

			actionMap := make(map[int]func())
			actionMap[1] = func() {
				defer wg.Done()
				foo.First(testPrintFirst)
			}

			actionMap[2] = func() {
				defer wg.Done()
				foo.Second(testPrintSecond)
			}

			actionMap[3] = func() {
				defer wg.Done()
				foo.Third(testPrintThird)
			}

			threadIds := BuildArray(tt.inputs)
			for _, id := range threadIds {
				wg.Add(1)
				go actionMap[id]()
			}

			wg.Wait()

			finalResult := strings.Join(actualOutput, "")
			if finalResult != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, finalResult)
			}
		})
	}
}
