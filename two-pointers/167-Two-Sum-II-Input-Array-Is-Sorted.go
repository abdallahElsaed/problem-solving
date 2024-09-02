package main 


func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	sum := 0

	for l < r {
		sum = numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}
		if sum > target {
			r -= 1
		}
		if sum < target {
			l += 1
		}
	}
	return nil

}
