package arraysandslices

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numsCol ...[]int) []int {
	var sums []int
	for _, nums := range numsCol {
		sums = append(sums, Sum(nums))
	}
	return sums
}

func SumAllTails(numsCol ...[]int) []int {
	var sums []int
	for _, nums := range numsCol {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
