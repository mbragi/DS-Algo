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