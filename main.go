package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, -2, 5}
	fmt.Printf("%d\n", KadaneAlgorithm(arr))
}

//KadaneAlgorithm test
func KadaneAlgorithm(arr []int) int {
	maxSoFar, maxHere := arr[0], arr[0]

	for i := 1; i < len(arr); i++ {

		if arr[i] < maxHere+arr[i] {
			maxHere = maxHere + arr[i]
		} else {
			maxHere = arr[i]
		}

		if maxSoFar < maxHere {
			maxSoFar = maxHere
		}
	}
	return maxSoFar
}
