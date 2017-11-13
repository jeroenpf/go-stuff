package main

import "fmt"

func main() {
	a := []int{1, 9, 12, 6, 4, 3, 2, 10, 30}
	fmt.Println("Before", a)
	fmt.Println("After", Mergesort(a))
}

func Mergesort(a []int) []int {

	if len(a) < 2 {
		return a
	}

	mid := len(a) / 2
	return merge(Mergesort(a[0:mid]), Mergesort(a[mid:]))
}

func merge(a []int, b []int) []int {

	out := []int{}

	for len(a) > 0 && len(b) > 0 {
		if a[0] > b[0] {
			out = append(out, b[0])
			b = b[1:]
		} else {
			out = append(out, a[0])
			a = a[1:]
		}
	}

	out = append(out, append(a, b...)...)

	return out
}
