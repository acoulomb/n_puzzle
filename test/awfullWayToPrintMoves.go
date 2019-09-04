package main

import (
	"fmt"
)

func iDidntWriteThat(tree ltreegraph, finalMove int) {
	var move []int
	move = append(move, finalMove)
	for tree.parent != nil {
		move = append([]int{tree.lastMove}, move...)
		tree = *tree.parent
	}
	move = append([]int{tree.lastMove}, move...)
	fmt.Println(move)
}
