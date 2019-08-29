package main

import (
	"fmt"
	"reflect"
)

type treegraph struct {
	hash   map[int]pos
	up     *treegraph
	down   *treegraph
	left   *treegraph
	right  *treegraph
	parent *treegraph
}

func keysByValue(hash map[int]pos, lat, lon int) int {
	for i, v := range hash {
		if v.Lat == lat && v.Lon == lon {
			return i
		}
	}
	return -1
}

func moveTree(r int, tree treegraph, EndHash map[int]pos) treegraph {
	child := new(treegraph)
	child.parent = &tree
	nMap := make(map[int]pos)
	switch r {
	case 1:
		toChange := keysByValue(tree.hash, tree.hash[0].Lat-1, tree.hash[0].Lon)
		if toChange == -1 {
			return *child
		}
		tree.up = child
		for i := 0; i < len(tree.hash); i++ {
			if i == toChange {
				nMap[0] = tree.hash[toChange]
			} else if i == 0 {
				nMap[toChange] = tree.hash[0]
			} else {
				nMap[i] = tree.hash[i]
			}
		}
		child.hash = nMap
		return *child
	case 2:
		toChange := keysByValue(tree.hash, tree.hash[0].Lat+1, tree.hash[0].Lon)
		if toChange == -1 {
			return *child
		}
		tree.down = child
		for i := 0; i < len(tree.hash); i++ {
			if i == toChange {
				nMap[0] = tree.hash[toChange]
			} else if i == 0 {
				nMap[toChange] = tree.hash[0]
			} else {
				nMap[i] = tree.hash[i]
			}
		}
		child.hash = nMap
		return *child
	case 3:
		toChange := keysByValue(tree.hash, tree.hash[0].Lat, tree.hash[0].Lon-1)
		if toChange == -1 {
			return *child
		}
		tree.left = child
		for i := 0; i < len(tree.hash); i++ {
			if i == toChange {
				nMap[0] = tree.hash[toChange]
			} else if i == 0 {
				nMap[toChange] = tree.hash[0]
			} else {
				nMap[i] = tree.hash[i]
			}
		}
		child.hash = nMap
		return *child
	case 4:
		toChange := keysByValue(tree.hash, tree.hash[0].Lat, tree.hash[0].Lon+1)
		if toChange == -1 {
			return *child
		}
		tree.right = child
		for i := 0; i < len(tree.hash); i++ {
			if i == toChange {
				nMap[0] = tree.hash[toChange]
			} else if i == 0 {
				nMap[toChange] = tree.hash[0]
			} else {
				nMap[i] = tree.hash[i]
			}
		}
		child.hash = nMap
		return *child
	}
	return *child
}

var dept int

func treeSearch(StockTree []map[int]pos, solution chan bool, tree treegraph, EndHash map[int]pos) {
	i := 0
	for deep := tree; deep.parent != nil; {
		i++
		deep = *deep.parent
	}
	if i > dept {
		dept = i
		fmt.Printf("depth is %d\n", i)
	}
	if tree.hash == nil {
		return
	}
	StockTree = append(StockTree, tree.hash)
	// fmt.Println(StockTree)
	if reflect.DeepEqual(tree.hash, EndHash) {
		solution <- true
		return
	}

	for i := 1; i <= 4; i++ {
		// fmt.Println(tree.hash)
		switch i {
		case 1:
			go func() {
				treeSearch(StockTree, solution, moveTree(1, tree, EndHash), EndHash)
			}()
		case 2:
			go func() {
				treeSearch(StockTree, solution, moveTree(2, tree, EndHash), EndHash)
			}()
		case 3:
			go func() {
				treeSearch(StockTree, solution, moveTree(3, tree, EndHash), EndHash)
			}()
		case 4:
			go func() {
				treeSearch(StockTree, solution, moveTree(4, tree, EndHash), EndHash)
			}()
		}
	}
	return
}

func aStar(base [][]int, EndHash map[int]pos) {
	solution := make(chan bool)
	tree := new(treegraph)
	tree.hash = HashMap(base)
	StockTree := make([]map[int]pos, 0)
	tree.parent = nil
	treeSearch(StockTree, solution, *tree, EndHash)
	<-solution
	return
}
