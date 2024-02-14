package main

import (
	"fmt"
)

func initial_positions(board [b_size][b_size]Object, positions ...Position) [b_size][b_size]Object {
	// расположение начальных организмов
	for _, position := range positions {
		board[position.x][position.y].present = 1
	}
	for i := 0; i < b_size; i++ {
		for j := 0; j < b_size; j++ {
			board[j][i].id = fmt.Sprintf("%d-%d", j, i)
		}
	}
	return board
}

func calc_borders(x int, y int, n int) Border {
	border := Border{
		x + 1,
		x - 1,
		y + 1,
		y - 1,
	}
	return border
}

func calc(board [b_size][b_size]Object) [b_size][b_size]Object {
	/*
		в пустой (мёртвой) клетке, с которой соседствуют три живые клетки, зарождается жизнь;
		если у живой клетки есть две или три живые соседки, то эта клетка продолжает жить;
		в противном случае (если живых соседей меньше двух или больше трёх) клетка умирает («от одиночества» или «от перенаселённости»).
	*/
	for i := 1; i < b_size-1; i++ {
		for j := 1; j < b_size-1; j++ {
			border := calc_borders(j, i, b_size)
			neighbor_1 := board[border.top][border.left].present
			neighbor_2 := board[border.top][i].present
			neighbor_3 := board[border.top][border.right].present
			neighbor_4 := board[j][border.right].present
			neighbor_5 := board[border.bottom][border.right].present
			neighbor_6 := board[border.bottom][i].present
			neighbor_7 := board[border.bottom][border.left].present
			neighbor_8 := board[j][border.left].present
			board[j][i].neighbors_count = neighbor_1 + neighbor_2 + neighbor_3 + neighbor_4 + neighbor_5 + neighbor_6 + neighbor_7 + neighbor_8

			if board[j][i].present == 0 {
				// клетка мертвая
				if board[j][i].neighbors_count >= 2 {
					board[j][i].future = 1
				}
			}
			if board[j][i].present == 1 {
				// клетка живая
				if board[j][i].neighbors_count > 3 || board[j][i].neighbors_count < 2 {
					board[j][i].future = 0
				}
			}
		}
	}
	for i := 1; i < b_size-1; i++ {
		for j := 1; j < b_size-1; j++ {
			board[j][i].update()
		}
	}
	return board
}

func print_board(board [b_size][b_size]Object) {
	// Печать текущего варианта доски
	for i := 0; i < b_size; i++ {
		for j := 0; j < b_size; j++ {
			fmt.Printf("%d ", board[j][i].present)
		}
		fmt.Println()
	}
	fmt.Println()
}

func board_status(board [b_size][b_size]Object) int {
	// число живых организмов
	var result int
	for i := 0; i < b_size; i++ {
		for j := 0; j < b_size; j++ {
			if board[j][i].present == 1 {
				result++
			}
		}
	}
	return result
}
