# Two Pointers problems

## Valid Palindrome
[problem link](https://leetcode.com/problems/valid-palindrome)

### What do you need in this problem?
* A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

### Solution:
* First, you need to filter all not alphabet or number characters from string and check if the word is palindrome or not.
#### Steps:
1. loop over the string 
2. filter the special characters from the string and put the chars or numbers on an array
3. make two pointers one from the first index and one on the last index
4. check if the element in the first not the element in the last index returns false.
5. after the loop is finished and does not return false you return true.

### code:
```
    func isPalindrome(s string) bool {
    	strOnly := make([]rune,0)
    	for _, char := range s { // you can loop over the string direct
    		if unicode.IsLetter(char) || unicode.IsDigit(char) {
    			strOnly = append(strOnly,unicode.ToLower(char)) // convert to string into small letter
    		}
    	}
    
    	last := len(strOnly) - 1 
    	for i := 0; i < len(strOnly)/2; i++ {
    		if strOnly[i] != strOnly[last] {
    			return false
    		}
    		last--
    	}
    	return true
    }
```