package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 5}
	result := findMajorityNumberAndFrequencyCount(arr)
	fmt.Println("majorityCount", result)
	
	arr2 := []int{1, 2, 1, 3, 6, 4, 5}
	fmt.Println("peakNumber", findPeakNumber(arr2))
	fmt.Println("peakNumberConcurrent", findPeakNumberConcurrent(arr2))
	
	linkedList	:= LinkedList{}
	linkedList.addNode(1)
	linkedList.addNode(2)
	printList(&linkedList)

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
