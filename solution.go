package solution

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}
		numMap[num] = i
	}

	return []int{}
}
