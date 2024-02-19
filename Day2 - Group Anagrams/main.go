package main

import "fmt"

func main() {
	res := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(res)
}

func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string) // Use a map to store anagrams grouped by their character count
	for _, s := range strs {
		SmallLettersCount := [26]int{}
		for _, c := range s {
			SmallLettersCount[int(c)-'a'] += 1
		}
		res[SmallLettersCount] = append(res[SmallLettersCount], s) // Add the string to the map
	}

	var anagrams [][]string
	for _, v := range res {
		anagrams = append(anagrams, v) // Append the anagrams to the result slice
	}
	return anagrams
}
