package main

type getNb struct {
	Nb  []int
	Pos []int
}

func (data *getNb) AppendPos(nb int, pos int) {
	data.Nb = append(data.Nb, nb)
	data.Pos = append(data.Pos, pos)
}

func mapEndState(endMap [][]int) (*getNb, *getNb) {
	// lat := make(map[int]int)
	lat := &getNb{[]int{}, []int{}}
	lon := &getNb{[]int{}, []int{}}
	for i := 0; i < len(endMap); i++ {
		for r := 0; r < len(endMap); r++ {
			lat.AppendPos(endMap[i][r], i)
			lon.AppendPos(endMap[i][r], r)
		}
	}
	return lat, lon
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
	for nb < totalnb {
		switch dir {
		case 1:
			for i := 0; i < move; i++ {
				lat++
				end_puzzle[lat][lon] = nb
				nb++
			}
			lon--
			dir = 2
		case 2:
			for i := 0; i < move; i++ {
				lon--
				end_puzzle[lat][lon] = nb
				nb++
			}
			lat--
			dir = 3
		case 3:
			for i := 0; i < move; i++ {
				lat--
				end_puzzle[lat][lon] = nb
				nb++
			}
			lon++
			dir = 4
		case 4:
			for i := 0; i < move; i++ {
				lon++
				end_puzzle[lat][lon] = nb
				nb++
			}
			lat++
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
