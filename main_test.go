package main

import "testing"

func TestKadaneAlgorithm(t *testing.T) {

	dataTable := []struct {
		arr                    []int
		expectedMaxSubSequence int
	}{
		{[]int{1, 2, 3, -2, 5}, 9},
		{[]int{1, 2, -3, -4, 2, 7, -2, 3}, 10},
		{[]int{-2, -3, -4, -2, -7, -2, -3, -11}, -2},
	}

	for _, data := range dataTable {
		actualOutput := KadaneAlgorithm(data.arr)
		if actualOutput == data.expectedMaxSubSequence {
			t.Log("SUCCESS")
		} else {
			t.Errorf("Failure, expected %d, actual %d", data.expectedMaxSubSequence, actualOutput)
		}
	}
}
