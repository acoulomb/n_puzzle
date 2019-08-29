package main

import (
	"fmt"

	"github.com/mgutz/ansi"
)

func main() {
	red := ansi.ColorFunc("red+")
	// if len(os.Args[1:]) < 1 {
	// 	fmt.Printf(red("▶ WARNING : Please insert THINGS as argument! \n"))
	// 	return
	// }
	starting_puzzle, err := parsing()
	if err != nil {
		fmt.Printf(red("▶ WARNING : %s \n"), err)
	}
	end_puzzle := get_end_state(starting_puzzle)
	_ = end_puzzle
	hashPos := HashMap(end_puzzle)
	// fmt.Println(manhattan_heuristic(starting_puzzle, hashPos))
	/*
	 **	 greedy won't work until rollback is fixed
	 */
	// GreedyResolve(starting_puzzle, hashPos)
	// GreedyResolve(starting_puzzle, hashPos)
	// aStarWoutHeuristic(starting_puzzle, hashPos)
	aStar(starting_puzzle, hashPos)
	// fmt.Println(starting_puzzle)
	// fmt.Println(hashPos)
	fmt.Println(end_puzzle, len(end_puzzle))
}
