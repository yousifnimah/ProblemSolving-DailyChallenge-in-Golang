package main

import (
	"fmt"
)

func main() {
	res := isAnagram("rat", "car")
	fmt.Println(res)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	HashMapS := make(map[string]int)
	HashMapT := make(map[string]int)
	for i := 0; i < len(s); i++ {
		HashMapS[string(s[i])], HashMapT[string(t[i])] = HashMapS[string(s[i])]+1, HashMapT[string(t[i])]+1
	}
	for ch, val := range HashMapS {
		if val != HashMapT[ch] {
			return false
		}
	}
	return true
}
