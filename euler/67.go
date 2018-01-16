// Euler 18 and 67
package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(maxPathSum(processTriangle()))
}

func maxPathSum(tree [][]int) int {
	for r := len(tree) - 2; r >= 0; r-- { // Start on the second last row
		for c := 0; c < len(tree[r]); c++ { // Go over each column in the tree
			tree[r][c] += int(math.Max(float64(tree[r+1][c]), float64(tree[r+1][c+1]))) // Add the maximum adjacent value to the current node
		}
	}
	return tree[0][0]
}

func processTriangle() [][]int {
	bContent, err := ioutil.ReadFile("resources/triangle.txt")
	if err != nil {
		panic(err)
	}

	triangle := string(bContent)
	tree := [][]int{}

	for _, line := range strings.Split(triangle, "\n") {
		if line == "" {
			continue
		}

		row := []int{}
		for _, val := range strings.Split(strings.TrimSpace(line), " ") {
			v, _ := strconv.Atoi(val)
			row = append(row, v)
		}
		tree = append(tree, row)
	}
	return tree
}
