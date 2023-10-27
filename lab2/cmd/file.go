package main

func shellSort(arr []int) []int {
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
	return arr
}

func findKthLargest(nums []int, k int) int {
	
	//Изначально использовал встроенную библиотеку sort
	//sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	//return nums[k-1]
	//но потом решил написать сортировку вручную

	nums = shellSort(nums)
	return nums[len(nums)-k]
}
