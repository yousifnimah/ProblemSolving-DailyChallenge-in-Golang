package main

import (
	"fmt"
	"sort"
)

func main() {
	res := isAnagram("tea", "eat")
	fmt.Println(res)
}

func isAnagram(s string, t string) bool {
	s = sortString(s)
	t = sortString(t)
	return s == t
}

func sortString(s string) string {
	t := []rune(s)
	sort.Slice(t, func(a, b int) bool {
		return t[a] < t[b]
	})
	return string(t)
}
