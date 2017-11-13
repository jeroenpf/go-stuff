package main_test

import (
	"fmt"
	"mergesort"
	"reflect"
	"testing"
)

type testpair struct {
	unsorted []int
	sorted   []int
}

var tests = []testpair{
	{[]int{1, 8, 2, 9, 3, 6, 4}, []int{1, 2, 3, 4, 6, 8, 9}},
	{[]int{10, 1, 49, 39, 5, 3, 100, 7}, []int{1, 3, 5, 7, 10, 39, 49, 100}},
}

func TestSorted(t *testing.T) {

	for _, pair := range tests {
		sorted := main.Mergesort(pair.unsorted)
		if reflect.DeepEqual(pair.sorted, sorted) == false {
			t.Error(fmt.Sprintf("Expected %v got %v", pair.sorted, sorted))
		}
	}
}
