package main

// Function to find peak number
func findPeakNumber(arr []int) []int {
	var result = []int{}
	if len(arr) == 0 {
		println("Array is empty")
		return result
	}

	for index, num := range arr {
		if index == 0 || index == len(arr)-1 {
			continue
		}
		if num > arr[index+1] && num > arr[index-1] {
			result = append(result, num)
		}
	}
	return result
}

func findPeakNumberConcurrent(arr []int) []int {
	var result []int
	if len(arr) == 0 {
		println("Array is empty")
		return result
	}

	ch := make(chan int)
	for index, num := range arr {
		go func(index, num int) {
			if index == 0 || index == len(arr)-1 {
				ch <- -1
				return
			}
			if num > arr[index+1] && num > arr[index-1] {
				ch <- num
			} else {
				ch <- -1
			}
		}(index, num)
	}

	for range arr {
		peak := <-ch
		if peak != -1 {
			result = append(result, peak)
		}
	}
	close(ch)
	return result
}
