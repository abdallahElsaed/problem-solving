# Hash Table and Array problems

## Contains Duplicate (1)
[problem link](https://leetcode.com/problems/group-anagrams/)

### What do you need in this problem?
* you have an array of words. Group the anagrams together.


### Solution:
* Sorting the characters of anagrams will always produce the same result. For example, both "eat" and "tea" will become "aet" when sorted. So we can use the map to use the sorted word to be a key of a map and the value is an array group of all the words that have the same sorted word.
[
    aet => [eat tea ate]
    ant => [tan nat]
    abt => [bat]
]

#### Steps:
1. loop over the strings array
2. sort the word 
3. check if this sorted word is in hashMap
    * if true:
        * add the word to the key value array 
    * if false:
        * create an array for this key
4. after finishing this loop make another loop over the hashMap
5. add the value for this map to new array to be a result

### code:

    func groupAnagrams(strs []string) [][]string {
	    hashMap := make(map[string][]string)
	    for _, str := range strs {
	    	sliceStr := strings.Split(str, "")
	    	sort.Strings(sliceStr)
	    	sortedStr := strings.Join(sliceStr, "")
	    	if arrayStrs, ok := hashMap[sortedStr]; ok {
	    		hashMap[sortedStr] = append(arrayStrs, str)
	    	} else {
	    		hashMap[sortedStr] = []string{str}
	    	}
	    }
	    resultArray := make([][]string, 0)
	    for _, array := range hashMap {
	    	resultArray = append(resultArray, array)
	    }
	    return resultArray
    }


