package main

import "fmt"

func main() {
	//! first problem (contains duplicate)
	// nums := []int{1, 2, 3}
	// fmt.Println(containsDuplicate(nums))

	//! seconde problem (Valid Anagram)
	// fmt.Println(isAnagram("nagaram" , "anagram"))

	//! third problem twoSum
	// nums := []int{3,4, 2, 1}
	// fmt.Println(twoSum(nums ,6))

	//! Fourth Problem group anagrams
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	
}
