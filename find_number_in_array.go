package main

import "fmt"

func main() {
	a := []int{1, 1000, 3000, 3001, 3278, 3942, 4000}
	b := []int{2000, 3001, 3942, 3950, 3960, 3970, 4000}

	offset := 0
	for _, it := range a {
		for of, bt := range b[offset:] {

			if bt > it {
				fmt.Printf("%s not found\n", it)
				break;
			}

			if(it == bt) {
				offset = of
				fmt.Printf("%s found\n", it)
				break;
			}

			offset = of
		}
	}
}
