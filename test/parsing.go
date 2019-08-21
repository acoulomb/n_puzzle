package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// func remove(slice [][]int, s int) [][]int {
// 	return append(slice[:s], slice[s+1:]...)
// }

func epur_tab(tab [][]int) [][]int {
	var new_tab [][]int
	for _, fields := range tab {
		if len(fields) > 1 {
			new_tab = append(new_tab, fields)
		}
	}
	return new_tab
}

func check_squared(clean_board [][]int) bool {
	for i, _ := range clean_board {
		if len(clean_board[0]) != len(clean_board[i]) {
			return false
		}
	}
	return true
}

func parsing() ([][]int, error) {
	var board [][]int
	stock_arg, err := ioutil.ReadAll(os.Stdin)
	new_arg := strings.Split(string(stock_arg), "\n")

	// gestion erreur TBD
	_ = err

	for _, line := range new_arg {
		var values []int
		fmt.Println(line)
		words := strings.Fields(line)
		for _, fields := range words {
			nb, err := strconv.Atoi(fields)
			if err != nil {
				if fields[0] == 35 {
					break
				}
				return nil, errors.New("Get THAT OUT -> " + fields)
			} else if err == nil {
				values = append(values, nb)
			}
		}
		board = append(board, values)
	}
	new_board := epur_tab(board)
	if !check_squared(new_board) {
		return new_board, errors.New("Not a square")
	}
	return epur_tab(board), nil
}
