package main

import (
	"fmt"
	"crypto/sha1"
)

type pos struct {
	Lat, Lon int
}

var m map[int]pos

func HashMap(endMap [][]int) map[int]pos {
	
	m = make(map[int]pos)
	for i, line := range endMap {
		for r, nb := range line {
			m[nb] = pos{
				i, r,
			}
		}
		_ = i
	}
	return m
}

func get_end_state(starting_puzzle [][]int) [][]int {
	end_puzzle := make([][]int, len(starting_puzzle))
	for i, line := range end_puzzle {
		for nb := 0; nb < len(end_puzzle); nb++ {
			line = append(line, -1)
		}
		end_puzzle[i] = line
	}
	nb := 1
	move := len(end_puzzle)
	totalnb := move * move
	var line []int
	for ; nb <= move; nb++ {
		line = append(line, nb)
	}
	move--
	end_puzzle[0] = line
	dir := 1
	itt := 0
	lat := 0
	lon := len(end_puzzle) - 1
	for nb <= totalnb {
		switch dir {
		case 1:
			for i := 0; i < move; i++ {
				lat++
				end_puzzle[lat][lon] = nb
				nb++
			}
			dir = 2
		case 2:
			for i := 0; i < move; i++ {
				lon--
				end_puzzle[lat][lon] = nb
				nb++
			}
			dir = 3
		case 3:
			for i := 0; i < move; i++ {
				lat--
				end_puzzle[lat][lon] = nb
				nb++
			}
			dir = 4
		case 4:
			for i := 0; i < move; i++ {
				lon++
				end_puzzle[lat][lon] = nb
				nb++
			}
			dir = 1
		}
		itt++
		if itt == 2 {
			itt = 0
			move--
		}
	}
	end_puzzle[lat][lon] = 0
	return end_puzzle
}
