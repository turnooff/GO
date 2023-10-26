package main

//import "sort"

func findKthLargest(nums []int, k int) int {
	//sort.Sort(sort.Reverse(sort.IntSlice(nums))) // sort.Ints(nums) можно использовать прямую сортировку
	//return nums[k-1] // nums[len(nums)-k] в случае прямой сортировки вывод такой
	var gap, i int
	for gap = len(nums) / 2; gap > 0; gap /= 2 {
		for i = gap; i < len(nums); i++ {
			var temp = nums[i]
			var j int
			for j = i; j >= gap && nums[j-gap] > temp; j -= gap {
				nums[j] = nums[j-gap]
			}
			nums[j] = temp
		}
	}
	return nums[len(nums)-k]
}
