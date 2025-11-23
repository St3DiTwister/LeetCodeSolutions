package _011_Container_With_Most_Water

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	for i < j {
		h := min(height[i], height[j])

		area := h * (j - i)
		if area > res {
			res = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return res
}
