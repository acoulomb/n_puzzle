package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func parsing(area string) [][]int {
	var board [][]int
	var values []int

	words := strings.Fields(area)
	for _, fields := range words {
		if fields[0] == 45 {
			board = append(board, values)
			values = nil
			continue
		}
		nb, err := strconv.Atoi(fields)
		values = append(values, nb)
		//TBD gestion err
		_ = err
	}
	board = append(board, values)
	return board
}

func resolve(area string) []int {
	startingPuzzle := parsing(area)
	endPuzzle := get_end_state(startingPuzzle)
	hashPos := HashMap(endPuzzle)
	return aStar(startingPuzzle, hashPos)
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		area := c.Query("area")
		c.String(http.StatusOK, "%d", resolve(area))
	})
	router.Run(":8080")
}
