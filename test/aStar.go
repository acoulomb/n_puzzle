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
	var nMap map[int]pos
	switch r {
	case 1:
		tree.up = child
		toChange := keysByValue(tree.hash, tree.hash[0].Lat-1, tree.hash[0].Lon)
		for k, v := range tree.hash {
			if k == toChange {
				nMap[0] = v
			} else if k == 0 {
				nMap[k] = tree.hash[0]
			} else {
				nMap[k] = v
			}
		}
		child.hash = nMap
		return *child
	case 2:
		tree.down = child
		toChange := keysByValue(tree.hash, tree.hash[0].Lat+1, tree.hash[0].Lon)
		for k, v := range tree.hash {
			if k == toChange {
				nMap[0] = v
			} else if k == 0 {
				nMap[k] = tree.hash[0]
			} else {
				nMap[k] = v
			}
		}
		child.hash = nMap
		return *child
	case 3:
		tree.left = child
		toChange := keysByValue(tree.hash, tree.hash[0].Lat, tree.hash[0].Lon-1)
		for k, v := range tree.hash {
			if k == toChange {
				nMap[0] = v
			} else if k == 0 {
				nMap[k] = tree.hash[0]
			} else {
				nMap[k] = v
			}
		}
		child.hash = nMap
		return *child
	case 4:
		tree.right = child
		toChange := keysByValue(tree.hash, tree.hash[0].Lat-1, tree.hash[0].Lon+1)
		for k, v := range tree.hash {
			if k == toChange {
				nMap[0] = v
			} else if k == 0 {
				nMap[k] = tree.hash[0]
			} else {
				nMap[k] = v
			}
		}
		child.hash = nMap
		return *child
	}
	return *child
}

func treeSearch(solution chan bool, tree treegraph, EndHash map[int]pos) {
	if tree.hash == nil {
		return
	}
	if reflect.DeepEqual(tree.hash, EndHash) {
		fmt.Println("patate")
		solution <- true
		return
	}
	for i := 1; i <= 4; i++ {
		switch i {
		case 1:
			go func() {
				treeSearch(solution, moveTree(1, tree, EndHash), EndHash)
			}()
		case 2:
			go func() {
				treeSearch(solution, moveTree(2, tree, EndHash), EndHash)
			}()
		case 3:
			go func() {
				treeSearch(solution, moveTree(3, tree, EndHash), EndHash)
			}()
		case 4:
			go func() {
				treeSearch(solution, moveTree(4, tree, EndHash), EndHash)
			}()
		}
	}
	return
}

func aStar(base [][]int, EndHash map[int]pos) {
	solution := make(chan bool)
	tree := new(treegraph)
	tree.hash = HashMap(base)
	tree.parent = nil
	treeSearch(solution, *tree, EndHash)
	<-solution
	return
}
