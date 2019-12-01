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

func insertInTree(gTree []ltreegraph, child ltreegraph) []ltreegraph {
	for i:= 0; i<= len(gTree); i++ {
		if i == len(gTree) {
			gTree = append(gTree, child)
			break
		}
		if child.heuristic < gTree[i].heuristic {
			gTree = append(gTree[:i], append([]ltreegraph{child}, gTree[i:]...)...)
			break
		}
	}
	return gTree
}

func arrayOfMove(gTree []ltreegraph, EndHash map[int]pos, StockTree []map[int]pos) []ltreegraph{
	tree := gTree[0]
	gTree = gTree[1:]
	var toChange int
	for r:= 0; r< 4; r++ {
		nMap := make(map[int]pos)
		child := new(ltreegraph)
		child.parent = &tree
		switch r {
		case 0:
			tree.up = child
			child.lastMove = 0
			toChange = keysByValue(tree.hash, tree.hash[0].Lat-1, tree.hash[0].Lon)
		case 1:
			tree.down = child
			child.lastMove = 1
			toChange = keysByValue(tree.hash, tree.hash[0].Lat+1, tree.hash[0].Lon)
		case 2:
			tree.left = child
			child.lastMove = 2
			toChange = keysByValue(tree.hash, tree.hash[0].Lat, tree.hash[0].Lon-1)
		case 3:
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
		if (checkPatern(StockTree, child.hash)) {
			gTree = insertInTree(gTree, *child)
		}
	}
	return gTree
}

func startTurn(StockTree []map[int]pos, gTree []ltreegraph, EndHash map[int]pos) {
	for ;gTree[0].heuristic != 0; {
		StockTree = append(StockTree, gTree[0].hash)
		gTree = arrayOfMove(gTree, EndHash, StockTree)
	}
	iDidntWriteThat(gTree[0])
	return
}


func aStar(base [][]int, EndHash map[int]pos) {
	var gTree []ltreegraph
	fmt.Println("beginaStar")
	tree := new(ltreegraph)
	tree.hash = HashMap(base)
	tree.heuristic = 42
	gTree = append(gTree, *tree)
	StockTree := make([]map[int]pos, 0)
	startTurn(StockTree, gTree, EndHash)
	return
}