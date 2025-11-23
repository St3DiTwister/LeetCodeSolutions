package _070_Climbing_Stairs

func climbStairs(n int) int {
	n += 1
	if n < 2 {
		return n
	}
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
