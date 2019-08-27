package main

import (
	"fmt"
	"reflect"
)

func checkPatern(tree []map[int]pos, pos map[int]pos) bool {
	// fmt.Println(tree)
	for i := 0; i < len(tree); i++ {
		if reflect.DeepEqual(tree[i], pos) {
			// fmt.Println(tree[i], pos)
			return false
		}
	}
	return true
}

func move(tree []map[int]pos, lastmove int, base [][]int, EndHash, hashBase map[int]pos) (int, [][]int) {
	var bestMove int
	var manHeuristic int
	var lowerHeuristic int
	mapCopy := make([][]int, len(base))
	bestCopy := make([][]int, len(base))
	for i := 1; i <= 4; i++ {
		for r := 0; r < len(base); r++ {
			mapCopy[r] = make([]int, len(base[r]))
			copy(mapCopy[r], base[r])
		}
		switch i {
		case 1:
			if hashBase[0].Lat-1 >= 0 {

				mapCopy[hashBase[0].Lat][hashBase[0].Lon] = mapCopy[hashBase[0].Lat-1][hashBase[0].Lon]
				mapCopy[hashBase[0].Lat-1][hashBase[0].Lon] = 0
				manHeuristic = manhattan_heuristic(mapCopy, EndHash)
				if (bestMove == 0 || manHeuristic < lowerHeuristic) && lastmove != i+1 && checkPatern(tree, HashMap(mapCopy)) {
					bestMove = i
					for r := 0; r < len(base); r++ {
						bestCopy[r] = make([]int, len(mapCopy[r]))
						copy(bestCopy[r], mapCopy[r])
					}
				}
			}
		case 2:
			if hashBase[0].Lat+1 < len(base) {

				mapCopy[hashBase[0].Lat][hashBase[0].Lon] = mapCopy[hashBase[0].Lat+1][hashBase[0].Lon]
				mapCopy[hashBase[0].Lat+1][hashBase[0].Lon] = 0
				manHeuristic = manhattan_heuristic(mapCopy, EndHash)

				if (bestMove == 0 || manHeuristic < lowerHeuristic) && lastmove != i-1 {
					bestMove = i
					for r := 0; r < len(base); r++ {
						bestCopy[r] = make([]int, len(mapCopy[r]))
						copy(bestCopy[r], mapCopy[r])
					}
				}
			}
		case 3:
			if hashBase[0].Lon-1 >= 0 {

				mapCopy[hashBase[0].Lat][hashBase[0].Lon] = mapCopy[hashBase[0].Lat][hashBase[0].Lon-1]
				mapCopy[hashBase[0].Lat][hashBase[0].Lon-1] = 0
				manHeuristic = manhattan_heuristic(mapCopy, EndHash)
				if (bestMove == 0 || manHeuristic < lowerHeuristic) && lastmove != i+1 {
					bestMove = i
					for r := 0; r < len(base); r++ {
						bestCopy[r] = make([]int, len(mapCopy[r]))
						copy(bestCopy[r], mapCopy[r])
					}
				}
			}
		case 4:
			if hashBase[0].Lon+1 < len(base) {

				mapCopy[hashBase[0].Lat][hashBase[0].Lon] = mapCopy[hashBase[0].Lat][hashBase[0].Lon+1]
				mapCopy[hashBase[0].Lat][hashBase[0].Lon+1] = 0
				manHeuristic = manhattan_heuristic(mapCopy, EndHash)
				if (bestMove == 0 || manHeuristic < lowerHeuristic) && lastmove != i-1 {
					bestMove = i
					for r := 0; r < len(base); r++ {
						bestCopy[r] = make([]int, len(mapCopy[r]))
						copy(bestCopy[r], mapCopy[r])
					}
				}
			}
		}
	}
	return bestMove, bestCopy
}

// need to rollback when bestmove==0

func GreedyResolve(base [][]int, EndHash map[int]pos) {
	bestMove := -1
	StockTree := make([]map[int]pos, 0)
	end_puzzle := make([][]int, len(base))
	for manhattan_heuristic(base, EndHash) != 0 {
		BaseHash := HashMap(base)
		StockTree = append(StockTree, BaseHash)
		fmt.Println(manhattan_heuristic(base, EndHash))
		bestMove, base = move(StockTree, bestMove, base, EndHash, BaseHash)
		if bestMove == 0 {
			fmt.Println("stuck in the middle")
			break
		}
		// fmt.Println(StockTree)
	}
	copy(end_puzzle, base)
	fmt.Println(EndHash)
	fmt.Println(end_puzzle, "end")
	return
}
