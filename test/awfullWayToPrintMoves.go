package main

import (
	"fmt"
)

func iDidntWriteThat(tree ltreegraph) {
	var move []int
	for tree.parent != nil {
		move = append([]int{tree.lastMove}, move...)
		tree = *tree.parent
	}
	move = append([]int{tree.lastMove}, move...)
	fmt.Println(move)
}
