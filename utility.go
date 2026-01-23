package solution

import (
	"strconv"
	"strings"
)

// BuildLinkedList creates a linked list from a string like "[2, 4, 5, 7]"
func BuildLinkedList(s string) *ListNode {
	// Remove brackets and spaces
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")

	// Split by comma
	parts := strings.Split(s, ",")
	if len(parts) == 0 || parts[0] == "" {
		return nil
	}

	// Create dummy node to simplify list building
	dummy := &ListNode{}
	current := dummy

	for _, part := range parts {
		// Parse the integer
		val, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			continue
		}

		// Create new node and append
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return dummy.Next
}

// SerialLinkedListToStr converts a linked list to a string representation like "[2, 4, 5, 7]"
func SerializeLinkedListToStr(head *ListNode) string {
	if head == nil {
		return "[]"
	}

	var result strings.Builder
	result.WriteString("[")

	current := head
	for current != nil {
		result.WriteString(strconv.Itoa(current.Val))
		if current.Next != nil {
			result.WriteString(",")
		}
		current = current.Next
	}

	result.WriteString("]")
	return result.String()
}
