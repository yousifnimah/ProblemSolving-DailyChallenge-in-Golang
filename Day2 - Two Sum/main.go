package main

import "fmt"

func main() {
	res := twoSum([]int{2, 11, 7, 3}, 9)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	HashMap := make(map[int]int)
	for i, n := range nums {
		comp := target - n
		val, ok := HashMap[comp]
		if ok {
			return []int{i, val}
		}
		HashMap[n] = i
	}
	return []int{0, 0}
}
