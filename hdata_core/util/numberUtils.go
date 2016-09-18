package util

/**
 * 获取 起始和结束 范围内的所有 数值
 *
 */
func GetRange(before, after int) []int {
	bigger := Max(before, after)
	smaller := Min(before, after)
	intList := make([]int, (bigger + 1 - smaller))
	for i := smaller; i <= bigger; i++ {
		intList = append(intList, i)
	}
	return intList
}
func Max(before, after int) int {
	if before > after {
		return before
	} else {
		return after
	}
}
func Min(before, after int) int {
	if before > after {
		return after
	} else {
		return before
	}
}
