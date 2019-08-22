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
	// mLat, mLon := mapEndState(end_puzzle)
	// fmt.Println(mLat, mLon)
	fmt.Println(starting_puzzle)
	fmt.Println(end_puzzle, len(end_puzzle))
	return
}
