package main

// Function to find peak number
func findPeakNumber(arr []int) []int  {
	var result	= []int{}
	if len(arr) == 0 {
		println("Array is empty")
		return result
	}

	for index, num := range arr {
		if index == 1 || num > arr[index+1] && num > arr[index-1] {
			result = append(result, num)
		}
	}
	return result
}
