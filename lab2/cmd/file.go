package main

func findKthLargest(nums []int, k int) int {
	//sort.Sort(sort.Reverse(sort.IntSlice(nums))) // sort.Ints(nums) можно использовать прямую сортировку
	//return nums[k-1] // nums[len(nums)-k] в случае прямой сортировки вывод такой
	helpArr := make([]int, k)
	for i := 0; i < k; i++ {
		//append(helpArr, nums[i])
		helpArr[i] = nums[i]
	}
	for i := k; i < len(nums); i++ {
		for j := 0; j < k; j++ {
			if nums[i] > helpArr[j] {
				helpArr[j] = nums[i]
			}
		}
	}
	min := helpArr[0]
	for i := 1; i < len(helpArr); i++ {
		if helpArr[i] < min {
			min = helpArr[i]
		}
	}
	return min
}
