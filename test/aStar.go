package main

import "fmt"

func move(lastmove int, base [][]int, EndHash, hashBase map[int]pos) (int, [][]int) {
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
				if (bestMove == 0 || manHeuristic < lowerHeuristic) && lastmove != i+1 {
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

func aStarResolve(base [][]int, EndHash map[int]pos) {
	bestMove := -1
	end_puzzle := make([][]int, len(base))
	for manhattan_heuristic(base, EndHash) != 0 {
		BaseHash := HashMap(base)
		fmt.Println(manhattan_heuristic(base, EndHash))
		bestMove, base = move(bestMove, base, EndHash, BaseHash)
		// fmt.Println(base)
	}
	copy(end_puzzle, base)
	fmt.Println(EndHash)
	fmt.Println(end_puzzle, "end")
	_ = bestMove
	return
}
