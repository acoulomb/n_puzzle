package main

func iDidntWriteThat(tree ltreegraph) []int {
	var move []int
	for tree.parent != nil {
		move = append([]int{tree.lastMove}, move...)
		tree = *tree.parent
	}
	move = append([]int{tree.lastMove}, move...)
	return move
}
