package main

// math.Abs in go takes float64 && it's penible to cast everything in int
func zouaveAbs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func manhattan_heuristic(base [][]int, EndHash map[int]pos) int {
	ManHeuristic := 0
	CurrentHash := HashMap(base)
	for i := 1; i < len(EndHash); i++ {
		ManHeuristic += zouaveAbs(CurrentHash[i].Lat-EndHash[i].Lat) + zouaveAbs(CurrentHash[i].Lon-EndHash[i].Lon)
	}
	return ManHeuristic
}
