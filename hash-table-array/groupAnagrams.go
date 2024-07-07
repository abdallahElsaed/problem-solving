package main

import (
	"sort"
	"strings"
)

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
