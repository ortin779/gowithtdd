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
