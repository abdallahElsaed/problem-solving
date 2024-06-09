package main

func containsDuplicate(nums []int) bool {
	checked_arr := make(map[int]bool)
	for _, num := range nums {
		if checked_arr[num] {
			return true
		}
		checked_arr[num] = true
	}
	return false
}
