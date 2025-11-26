package _046_Permutations

func fact(n int) int {
	if n <= 1 {
		return n
	}
	return n * fact(n-1)
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	result := make([][]int, 0, fact(len(nums)))
	for i := 0; i < len(nums); i++ {
		first := nums[i]
		rest := append(append([]int{}, nums[:i]...), nums[i+1:]...)
		perms := permute(rest)
		for _, p := range perms {
			newPerm := append([]int{first}, p...)
			result = append(result, newPerm)
		}
	}

	return result
}
