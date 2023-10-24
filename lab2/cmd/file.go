package main

import (
	"fmt"
	"sort"
)

//Example of tested file
func findKthLargest(nums []int, k int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // sort.Ints(nums) можно использовать прямую сортировку
	fmt.Println(nums[k-1]) // fmt.Println(nums[len(nums)-k]) в случае прямой сортировки вывод такой
	return 0

}
