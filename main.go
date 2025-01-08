package main

import (
	// "fmt"
	"math"
)

// Function to find majority number and its count
func findMajorityNumberAndFrequencyCount(arr []int) map[string]interface{} {
	if len(arr) == 0 {
		return map[string]interface{}{"majority": nil, "count": 0}
	}

	freqMap := make(map[int]int)
	n := len(arr)
	for _, num := range arr {
		if freqMap[num] > int(math.Floor(float64(n)/2)) {
			return map[string]interface{}{"majority": num, "count": freqMap[num]}
		}
		freqMap[num]++
	}
	return map[string]interface{}{"majority": nil, "count": 0}
}

func main() {
	// arr := []int{1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 5}
	// result := findMajorityNumberAndFrequencyCount(arr)
	// fmt.Println("majorityCount", result)
	
	// arr2 := []int{1, 2, 1, 3, 6, 4, 5}
	// fmt.Println("peakNumber", findPeakNumber(arr2))
	// fmt.Println("peakNumberConcurrent", findPeakNumberConcurrent(arr2))
	
	// linkedList	:= LinkedList{}
	// linkedList.addNode(1)
	// linkedList.addNode(2)
	// printList(&linkedList)

	bt := &Tree{0, nil, nil}
	bt.Insert(2)
	bt.Insert(40)
	bt.Insert(3)
	bt.Insert(48)
	bt.Insert(8)
	bt.Insert(9)
	// bt.BFS()
	bt.DFS()
}
