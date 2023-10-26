package main

//import "sort"

func findKthLargest(nums []int, k int) int {
	//sort.Sort(sort.Reverse(sort.IntSlice(nums))) // sort.Ints(nums) можно использовать прямую сортировку
	//return nums[k-1] // nums[len(nums)-k] в случае прямой сортировки вывод такой
	var i, j, key int
	for i = 1; i < len(nums); i++ {
		key = nums[i]
		j = i - 1

		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j = j - 1
		}
		nums[j+1] = key
	}
	return nums[len(nums)-k]
}
