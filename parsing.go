package main

import (
	"errors"
	"fmt"
	"os"
)

func parsing() ([]int, error) {
	var test []int
	stock_arg := os.Args[1:]
	if len(stock_arg) != len(stock_arg[0]) {
		return nil, errors.New("Not a square")
	}
	fmt.Println(stock_arg, len(stock_arg), len(stock_arg[0]), stock_arg[0])
	return test, nil
}
