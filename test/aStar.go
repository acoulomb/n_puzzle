package main

import (
	"fmt"
)

type ltreegraph struct {
	lastMove  int
	heuristic int
	hash      map[int]pos
	up        *ltreegraph
	down      *ltreegraph
	left      *ltreegraph
	right     *ltreegraph
	parent    *ltreegraph
}

func arrayOfMoves(tree ltreegraph, EndHash map[int]pos) []ltreegraph {
	var possibleMove []ltreegraph
	var toChange int
	for r := 0; r < 4; r++ {
		nMap := make(map[int]pos)
		child := new(ltreegraph)
		if r == 0 {
			child.lastMove = 0
			toChange = keysByValue(tree.hash, tree.hash[0].Lat-1, tree.hash[0].Lon)
		} else if r == 1 {
			child.lastMove = 1
			toChange = keysByValue(tree.hash, tree.hash[0].Lat+1, tree.hash[0].Lon)
		} else if r == 2 {
			child.lastMove = 2
			toChange = keysByValue(tree.hash, tree.hash[0].Lat, tree.hash[0].Lon-1)
		} else {
			child.lastMove = 3
			toChange = keysByValue(tree.hash, tree.hash[0].Lat, tree.hash[0].Lon+1)
		}
		if toChange == -1 {
			continue
		}
		for i := 0; i < len(tree.hash); i++ {
			if i == toChange {
				nMap[0] = tree.hash[toChange]
			} else if i == 0 {
				nMap[toChange] = tree.hash[0]
			} else {
				nMap[i] = tree.hash[i]
			}
		}
		tree.up = child
		child.heuristic = manWithHash(nMap, EndHash)
		child.hash = nMap
		possibleMove = append(possibleMove, *child)
	}
	return possibleMove
}

func startTurn(StockTree []map[int]pos, tree ltreegraph, EndHash map[int]pos) {
	possibleMove := arrayOfMoves(tree, EndHash)
	for _, ltree := range possibleMove {
		fmt.Println("test", ltree)
	}
	// //pop element of slice
	// r := 0
	// possibleMove[r] = possibleMove[len(possibleMove)-1]
	// possibleMove = possibleMove[:len(possibleMove)-1]
	return
}

func aStar(base [][]int, EndHash map[int]pos) {
	tree := new(ltreegraph)
	tree.hash = HashMap(base)
	StockTree := make([]map[int]pos, 0)
	startTurn(StockTree, *tree, EndHash)
	return
}
