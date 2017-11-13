package main

import "fmt"

func main() {
	a := []int{30, 1, 6, 29, 34, 8, 2, 5, 10, 4, 12, 3}
	fmt.Println("Before", a)
	quicksort(a)
	fmt.Println("After", a)
}

func quicksort(a []int) {
	if len(a) <= 1 {
		return
	}

	index := partition(a)
	quicksort(a[:index])
	quicksort(a[index+1:])
}

func partition(a []int) int {
	index := 0
	pivot := a[len(a)-1]
	for i := 0; i < len(a)-1; i++ {
		if a[i] <= pivot {
			a[i], a[index] = a[index], a[i]
			index++
		}
	}

	a[index], a[len(a)-1] = a[len(a)-1], a[index]
	return index
}
