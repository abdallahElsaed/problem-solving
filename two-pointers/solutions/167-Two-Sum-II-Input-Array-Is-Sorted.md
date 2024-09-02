# 167. Two Sum II - Input Array Is Sorted

## Valid Palindrome
[problem link](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted)

### What do you need in this problem?
* You have an sorted array and number. you want to get the two indexed number to sum those to get the target.

### Solution:
* you use two pointer to solve this problem. one on the left side and other in right side. and you move two pointer based on condition.
#### Steps:
1. declare three variable left, right, sum all is int
2. loop over the array of numbers while the left less than the right to not overlap or cross each other
3. check if the sum equal the target return the left and right but add 1 to those to fit the requirement of the problem.
4. check if the sum greater then the target less the right by 1 to less the sum.
5. check if the sum less then the target add 1 to the left to greet the sum.

### code:
```
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
```