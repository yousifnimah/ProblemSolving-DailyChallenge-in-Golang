package main

import "fmt"

func main() {
	res := productExceptSelf([]int{1, 2, 3, 4})
	fmt.Println(res)
}

func productExceptSelf(nums []int) []int {
	prefix := 1
	postfix := 1
	n := len(nums)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	for i := n - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}
	return res
}
