package Problem0011

func maxArea(height []int) int {
	// 从两端开始寻找，至少保证了宽度是最大值
	i, j := 0, len(height)-1
	max := 0

	for i < j {
		a, b := height[i], height[j]
		h := min(a, b)

		area := h * (j - i)
		if max < area {
			max = area
		}

		// 朝着area具有变大的可能性方向变化。
		if a < b {
			i++
		} else {
			j--
		}
	}

	return max
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func MaxAreaYy(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		width := right - left
		currentArea := min2(height[left], height[right]) * width

		if currentArea > maxArea {
			maxArea = currentArea
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return maxArea
}

func min2(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
