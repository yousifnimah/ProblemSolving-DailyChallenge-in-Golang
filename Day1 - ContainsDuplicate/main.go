package main

import "fmt"

func main() {
	res := containsDuplicate([]int{1, 2, 3, 1})
	fmt.Println(res)
}

func containsDuplicate(nums []int) bool {
	hashSet := make(map[int]int)
	for i, num := range nums {
		_, ok := hashSet[num]
		if ok {
			return true
		}
		hashSet[num] = i
	}
	return false
}
