package main

func findKthLargest(nums []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // sort.Ints(nums) можно использовать прямую сортировку
	return nums[k-1] // nums[len(nums)-k] в случае прямой сортировки вывод такой
}
