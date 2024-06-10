package main
import(
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	// convert string to slice
	slice1 := strings.Split(s, "")
	slice2 := strings.Split(t, "")
	//Sort this slice
	sort.Strings(slice1)
	sort.Strings(slice2)
	//Convert the slice to a string
	s = strings.Join(slice1, "")
	t = strings.Join(slice2, "")
	if s == t {
		return true
	} else {
		return false
	}
}