package _136_Single_Number

func singleNumber(nums []int) int {
	nums_arr := make(map[int]int, len(nums))

	for _, num := range nums {
		nums_arr[num] += 1
	}
	for num, count := range nums_arr {
		if count == 1 {
			return num
		}
	}
	return 0
}
