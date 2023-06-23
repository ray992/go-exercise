package main

/**
	直方图最大矩形面积
**/
func largestRectangleArea(heights []int) int {
	max, left, right := 0, 0, 0
	for ii := 0; ii < len(heights); ii++ {
		left = ii - 1
		right = ii + 1
		if len(heights)*heights[ii] > max {
			for left >= 0 && heights[left] >= heights[ii] {
				left--
			}
			for right <= len(heights)-1 && heights[right] >= heights[ii] {
				right++
			}
			curArea := (right - left - 1) * heights[ii]
			if curArea > max {
				max = curArea
			}
		}
	}
	return max
}
