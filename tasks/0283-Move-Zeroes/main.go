package _283_Move_Zeroes

func moveZeroes(nums []int) []int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			i++
		}
	}

	return nums
}
