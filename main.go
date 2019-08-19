package main

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

func main() {
	red := ansi.ColorFunc("red+")
	if len(os.Args[1:]) < 1 {
		fmt.Printf(red("▶ WARNING : Please insert THINGS as argument! \n"))
		return
	}
	starting_puzzle, err := parsing()
	if err != nil {
		fmt.Printf(red("▶ WARNING : %s \n"), err)
	}
	_ = starting_puzzle
	return
}
