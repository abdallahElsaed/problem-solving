# Hash Table and Array problems

## Contains Duplicate (1)
[problem link](https://leetcode.com/problems/contains-duplicate/submissions/)

### What do you need in this problem?
* you have an array you need to know if the array has duplicate numbers or not.

### Solution:
* you have two ways to solve this problem:
    1. iteration over the array and compare every single number on the array and make nested loop
        ```go
            func containsDuplicate(nums []int) bool {
                for i := 0; i < len(nums); i++ {
                    for j := i + 1; j < len(nums); j++ { // Start j from i + 1
                        if nums[i] == nums[j] {
                            return true // Return true if a duplicate is found
                        }
                    }
                }
                return false // Return false if no duplicates are found
            }
        ```
        > In this solution the time complexity is O(n2)
        #### Steps:
        * you will loop over the array
        * check if the first num equals any number on the array by make another loop
        * this technique called (porte force)

    2. second way to use a hash map:
        ```go
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
        ```
        > In this solution the time complexity is O(n)
        #### Steps:
        * we will loop over an array
        * make hash map (hash table data structure) key-value 
        * the key ie number and value is true // to return true when the key is exist 
        * check if the key is exist
            * if exist => return true
            * if not => add this key to map
        * after the loop is finished and does not return true this means the array is not duplicated. *and return false*
    
