# Hash Table and Array problems

## Contains Duplicate (1)
[problem link](https://leetcode.com/problems/two-sum/)

### What do you need in this problem?
* You have an array and an integer number. you need to return indices of the two numbers such that they add up to a target.

### Solution:
* As you iterate through the array, for each element, calculate the complement by subtracting the element from the target. Check if the complement is already in your hash map. If it is, you've found your two numbers. If it’s not, add the current number and its index to the hash map. This way, you can quickly check if you've seen the complement before without iterating over the array multiple times.
#### Steps:
1. create hash map
2. iterate over the array and calculate the complement by subtracting the element from the target.
3. if the complement inside the map
    * return the array containing the two number
4. if not 
    * add the element to the map


### code:

    ```go
        func twoSum(nums []int, target  int) []int {
	        numMap := make(map[int]int)
            for i, num := range nums {
                complement := target - num
                if index, ok := numMap[complement]; ok {
                    return []int{index, i}
                }
                numMap[num] = i
            }
            return nil
        }
    ```

