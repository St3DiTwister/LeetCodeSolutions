package _278_First_Bad_Version

func firstBadVersion(n int, isBadVersion func(int) bool) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}
