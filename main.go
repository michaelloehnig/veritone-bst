package main

import (
	"fmt"
	"veritone-bst/bst"
)

func main() {
	tree := bst.FromSlice([]int{12, 11, 90, 82, 7, 9})
	values, depth, err := tree.DeepestNodes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("deepest: %v, depth: %v\n", values, depth)
}
