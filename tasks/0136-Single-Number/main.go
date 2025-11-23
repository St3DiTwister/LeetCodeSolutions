package _136_Single_Number

func singleNumber(nums []int) int {
	numsArr := make(map[int]int, len(nums))

	for _, num := range nums {
		numsArr[num] += 1
	}
	for num, count := range numsArr {
		if count == 1 {
			return num
		}
	}
	return 0
}
