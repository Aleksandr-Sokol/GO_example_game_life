package main

import (
	"fmt"
)

const b_size = 10
const time_steps = 100

func main() {
	fmt.Println("START")

	var zero_board [b_size][b_size]Object
	board := initial_positions(zero_board,
		Position{1, 1},
		Position{1, 2},
		Position{2, 2},
		Position{3, 3},
	)

	max_status := -1
	for i := 1; i < time_steps; i++ {
		status := board_status(board)
		fmt.Println(status)
		if max_status < status {
			max_status = status
		}
		print_board(board)
		board = calc(board)
	}

	fmt.Println("STOP", float64(max_status)/(b_size*b_size)*100)
}
