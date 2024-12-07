package leetcode

// nums 바꿔야하고 몇개인지 k return
func removeElement(nums []int, val int) int {
	// val 나오면 한칸씩 땡기다가
	// 또 나오면 두칸씩 땡기다가
	// 마지막에 뺀 수 만큼 짜르기
	rk := 0

	for i, v := range nums {
		nums[i-rk] = nums[i]
		if v == val {
			rk++
		}
	}

	return len(nums) - rk
}
