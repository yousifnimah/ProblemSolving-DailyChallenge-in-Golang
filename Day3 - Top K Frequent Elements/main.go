package main

import "fmt"

func main() {
	res := topKFrequent([]int{1, 1, 1, 5, 5, 3}, 2)
	fmt.Println(res)
}

func topKFrequent(nums []int, k int) []int {
	HashMap := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		HashMap[nums[i]]++
	}
	bucket := make([][]int, n+1) // we create slice of slices of integer
	for n, frq := range HashMap {
		bucket[frq] = append(bucket[frq], n)
	}

	var res []int
	for j := len(bucket) - 1; j >= 0; j-- {
		size := len(bucket[j]) + len(res)
		if size <= k {
			res = append(res, bucket[j]...)
		}
	}
	return res
}
