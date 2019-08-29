package main

import (
	"fmt"
	"math"
)

func printTab(hash map[int]pos) {
	for i := 0; i < int(math.Sqrt(float64(len(hash)))); i++ {
		for r := 0; r < int(math.Sqrt(float64(len(hash)))); r++ {
			fmt.Printf("%d ", keysByValue(hash, i, r))
		}
		fmt.Printf("\n")
	}
}
