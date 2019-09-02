package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
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
		child.parent = &tree
		if r == 0 {
			tree.up = child
			child.lastMove = 0
			toChange = keysByValue(tree.hash, tree.hash[0].Lat-1, tree.hash[0].Lon)
		} else if r == 1 {
			tree.down = child
			child.lastMove = 1
			toChange = keysByValue(tree.hash, tree.hash[0].Lat+1, tree.hash[0].Lon)
		} else if r == 2 {
			tree.left = child
			child.lastMove = 2
			toChange = keysByValue(tree.hash, tree.hash[0].Lat, tree.hash[0].Lon-1)
		} else {
			tree.right = child
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
		child.heuristic = manWithHash(nMap, EndHash)
		child.hash = nMap
		possibleMove = append(possibleMove, *child)
	}
	return possibleMove
}

func sortArrayByHeuristic(possibleMove []ltreegraph) []ltreegraph {
	sort.Slice(possibleMove, func(i, j int) bool {
		switch strings.Compare(string(possibleMove[i].heuristic), string(possibleMove[j].heuristic)) {
		case -1:
			return true
		case 1:
			return false
		}
		return possibleMove[i].heuristic > possibleMove[j].heuristic
	})
	return possibleMove
}

// var deep int

func startTurn(StockTree []map[int]pos, tree ltreegraph, EndHash map[int]pos) {
	possibleMove := arrayOfMoves(tree, EndHash)
	for i := 0; i < len(possibleMove); i++ {
		if reflect.DeepEqual(tree.hash, EndHash) {
			os.Exit(1)
		}
	}
	fmt.Println(possibleMove[0].parent)
	// start := 0
	// copymove := possibleMove[0]
	// for copymove.parent != nil {
	// 	copymove = *copymove.parent
	// 	start++
	// }
	// if start > deep {
	// 	fmt.Printf("dept is %d\n", start)
	// 	deep = start
	// }
	// if loop == 10 {
	// 	os.Exit(-1)
	// }
	// loop++
	// for _, ltree := range possibleMove {
	// 	fmt.Println("test", ltree)
	// }
	possibleMove = sortArrayByHeuristic(possibleMove)
	for len(possibleMove) > 0 {
		// fmt.Println(possibleMove)
		startTurn(StockTree, possibleMove[0], EndHash)
		possibleMove[0] = possibleMove[len(possibleMove)-1]
		possibleMove = possibleMove[:len(possibleMove)-1]
		possibleMove = sortArrayByHeuristic(possibleMove)
	}
	// //pop element of slice
	// r := 0
	return
}

func aStar(base [][]int, EndHash map[int]pos) {
	tree := new(ltreegraph)
	tree.hash = HashMap(base)
	StockTree := make([]map[int]pos, 0)
	startTurn(StockTree, *tree, EndHash)
	return
}
