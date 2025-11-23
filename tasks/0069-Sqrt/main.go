package _069_Sqrt

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	low := 0
	high := x

	var mid int

	for low <= high {
		mid = (low + high) / 2
		res := mid * mid

		if res == x {
			return mid
		} else if res > x {
			high = mid - 1
		} else if res < x {
			low = mid + 1
		}
	}
	return high
}
