package _128_最长连续序列

func ContinuousSequence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var hashSet = make(map[int]bool)
	for _, num := range nums {
		hashSet[num] = true
	}

	var maxLength int
	for _, num := range nums {
		if !hashSet[num-1] {
			currentNum := num
			currentLength := 1
			for hashSet[num+1] {
				currentNum++
				currentLength++
			}
			if maxLength < currentLength {
				maxLength = currentLength
			}
		}
	}
	return maxLength
}
