package leetcode

func majorityElement(nums []int) int {
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++
	}

	majorityCnt := len(nums) / 2
	for n, v := range numMap {
		if v > majorityCnt {
			return n
		}
	}
	return 0
}
